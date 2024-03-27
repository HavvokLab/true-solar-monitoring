package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/olivere/elastic/v7"
)

type SolarRepo interface {
	BulkIndex(index string, docs []interface{}) error
	GetPlantDailyProduction(start, end *time.Time) ([]*elastic.AggregationBucketCompositeItem, error)
	GetPlantMonthlyProduction(start, end *time.Time) ([]*elastic.AggregationBucketCompositeItem, error)
	UpsertSiteStation(docs []model.SiteItem) error
	GetPerformanceLow(duration int, efficiencyFactor float64, focusHour int, thresholdPct float64) ([]*elastic.AggregationBucketCompositeItem, error)
	GetPerformanceByDate(date *time.Time, efficiencyFactor float64, focusHour int, thresholdPct float64) ([]*elastic.AggregationBucketCompositeItem, error)
	GetSumPerformanceLow(duration int) ([]*elastic.AggregationBucketCompositeItem, error)
	GetUniquePlantByIndex(index string) ([]*elastic.AggregationBucketKeyItem, error)
	UpdateOwnerToIndex(index, owner string) error
	DeleteIndex(index string) error
	DeleteDocumentInIndex(index, id string) error
}

type solarRepo struct {
	elastic *elastic.Client
	config  config.ElasticsearchConfig
}

func NewSolarRepo(elastic *elastic.Client) SolarRepo {
	conf := config.GetConfig().Elastic
	return &solarRepo{
		elastic: elastic,
		config:  conf,
	}
}

func (r *solarRepo) searchIndex() *elastic.SearchService {
	index := fmt.Sprintf("%v*", r.config.SolarIndex)
	return r.elastic.Search(index)
}

func (r *solarRepo) createIndexIfNotExist(index string) error {
	ctx := context.Background()
	if exist, err := r.elastic.IndexExists(index).Do(ctx); err != nil {
		if !exist {
			result, err := r.elastic.CreateIndex(index).Do(ctx)
			if err != nil {
				return err
			}

			if !result.Acknowledged {
				return errors.New("elasticsearch did not acknowledge")
			}
		}
	}

	return nil
}

// |=> Implementation
func (r *solarRepo) BulkIndex(index string, docs []interface{}) error {
	if err := r.createIndexIfNotExist(index); err != nil {
		return err
	}

	bulk := r.elastic.Bulk()
	for _, doc := range docs {
		bulk.Add(elastic.NewBulkIndexRequest().Index(index).Doc(doc))
	}

	ctx := context.Background()
	if _, err := bulk.Do(ctx); err != nil {
		return err
	}

	return nil
}

func (r *solarRepo) UpsertSiteStation(docs []model.SiteItem) error {
	index := config.GetConfig().Elastic.SiteStationIndex
	err := r.createIndexIfNotExist(index)
	if err != nil {
		return err
	}

	bulk := r.elastic.Bulk()
	for _, doc := range docs {
		bulk.Add(elastic.NewBulkUpdateRequest().Index(index).Id(doc.SiteID).Doc(doc).DocAsUpsert(true))
	}

	_, err = bulk.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *solarRepo) GetPlantDailyProduction(start, end *time.Time) ([]*elastic.AggregationBucketCompositeItem, error) {
	ctx := context.Background()
	items := make([]*elastic.AggregationBucketCompositeItem, 0)

	// create [composite aggregation]
	compositeAggregation := elastic.NewCompositeAggregation().
		Size(10000).
		Sources(
			elastic.NewCompositeAggregationDateHistogramValuesSource("date").Field("@timestamp").CalendarInterval("day").Format("yyyy-MM-dd").TimeZone("Asia/Bangkok"),
			elastic.NewCompositeAggregationTermsValuesSource("vendor_type").Field("vendor_type.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("area").Field("area.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("name").Field("name.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("owner").Field("owner.keyword"),
		)

	// assign [max_aggregation] into composite aggregation
	compositeAggregation = compositeAggregation.
		SubAggregation("installed_capacity", elastic.NewMaxAggregation().Field("installed_capacity")).
		SubAggregation("monthly_production", elastic.NewMaxAggregation().Field("monthly_production")).
		SubAggregation("daily_production", elastic.NewMaxAggregation().Field("daily_production")).
		SubAggregation(
			"lat",
			elastic.NewMaxAggregation().
				Script(
					elastic.NewScript("if (doc['location'].size() == 0) {return 0} else {doc['location'].lat}").
						Lang("painless"),
				),
		).
		SubAggregation(
			"long",
			elastic.NewMaxAggregation().
				Script(
					elastic.NewScript("if (doc['location'].size() == 0) {return 0} else {doc['location'].lon}").
						Lang("painless"),
				),
		)

	// assign [bucket_script_aggregation] into composite aggregation
	// |=> target
	const targetScript = "if (params.installed_capacity == 0 || params.daily_production == 0 ) { return 0 } else { (params.installed_capacity*5*0.8) }"

	compositeAggregation = compositeAggregation.SubAggregation(
		"target",
		elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"installed_capacity": "installed_capacity", "daily_production": "daily_production"}).
			Script(elastic.NewScript(targetScript)),
	)

	// |=> production_to_target
	const productionToTargetScript = "if (params.installed_capacity == 0 || params.daily_production == 0 ) { return 0 } else { (params.daily_production/(params.installed_capacity*5*0.8))*100 }"

	compositeAggregation = compositeAggregation.SubAggregation(
		"production_to_target",
		elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"installed_capacity": "installed_capacity", "daily_production": "daily_production"}).
			Script(elastic.NewScript(productionToTargetScript)),
	)

	query := r.searchIndex().
		Size(0).
		Query(
			elastic.NewBoolQuery().Must(
				elastic.NewTermQuery("data_type.keyword", constant.DATA_TYPE_PLANT),
				elastic.NewRangeQuery("@timestamp").Gte(start).Lt(end).TimeZone("Asia/Bangkok"),
				elastic.NewTermsQuery("vendor_type.keyword",
					strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
					strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
					strings.ToUpper(constant.VENDOR_TYPE_INVT),
					strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
				),
			),
		).Aggregation("production", compositeAggregation)

	result, err := query.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("cannot get result")
	}

	if result.Aggregations == nil {
		return nil, errors.New("cannot get result aggregation")
	}

	productions, found := result.Aggregations.Composite("production")
	if !found {
		return nil, errors.New("cannot get result composite performance alarm")
	}

	items = append(items, productions.Buckets...)
	return items, nil
}

func (r *solarRepo) GetPlantMonthlyProduction(start, end *time.Time) ([]*elastic.AggregationBucketCompositeItem, error) {
	ctx := context.Background()
	items := make([]*elastic.AggregationBucketCompositeItem, 0)

	// create [composite aggregation]
	compositeAggregation := elastic.NewCompositeAggregation().
		Size(10000).
		Sources(
			elastic.NewCompositeAggregationDateHistogramValuesSource("date").Field("@timestamp").CalendarInterval("month").Format("yyyy-MM-dd").TimeZone("Asia/Bangkok"),
			elastic.NewCompositeAggregationTermsValuesSource("vendor_type").Field("vendor_type.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("area").Field("area.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("name").Field("name.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("owner").Field("owner.keyword"),
		)

	// assign [max_aggregation] into composite aggregation
	compositeAggregation = compositeAggregation.
		SubAggregation("installed_capacity", elastic.NewMaxAggregation().Field("installed_capacity")).
		SubAggregation("monthly_production", elastic.NewMaxAggregation().Field("monthly_production")).
		SubAggregation("daily_production", elastic.NewMaxAggregation().Field("daily_production")).
		SubAggregation(
			"lat",
			elastic.NewMaxAggregation().
				Script(
					elastic.NewScript("if (doc['location'].size() == 0) {return 0} else {doc['location'].lat}").
						Lang("painless"),
				),
		).
		SubAggregation(
			"long",
			elastic.NewMaxAggregation().
				Script(
					elastic.NewScript("if (doc['location'].size() == 0) {return 0} else {doc['location'].lon}").
						Lang("painless"),
				),
		)

	// assign [bucket_script_aggregation] into composite aggregation
	// |=> target
	const targetScript = "if (params.installed_capacity == 0 || params.daily_production == 0 ) { return 0 } else { (params.installed_capacity*5*0.8*31) }"

	compositeAggregation = compositeAggregation.SubAggregation(
		"target",
		elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"installed_capacity": "installed_capacity", "daily_production": "daily_production"}).
			Script(elastic.NewScript(targetScript)),
	)

	// |=> production_to_target
	const productionToTargetScript = "if (params.installed_capacity == 0 || params.monthly_production == 0 ) { return 0 } else { (params.monthly_production/(params.installed_capacity*5*0.8*31))*100 }"

	compositeAggregation = compositeAggregation.SubAggregation(
		"production_to_target",
		elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"installed_capacity": "installed_capacity", "monthly_production": "monthly_production"}).
			Script(elastic.NewScript(productionToTargetScript)),
	)

	query := r.searchIndex().
		Size(0).
		Query(
			elastic.NewBoolQuery().Must(
				elastic.NewTermQuery("data_type.keyword", constant.DATA_TYPE_PLANT),
				elastic.NewRangeQuery("@timestamp").Gte(start).Lt(end).TimeZone("Asia/Bangkok"),
				elastic.NewTermsQuery("vendor_type.keyword",
					strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
					strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
					strings.ToUpper(constant.VENDOR_TYPE_INVT),
					strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
				),
			),
		).Aggregation("production", compositeAggregation)

	result, err := query.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("cannot get result")
	}

	if result.Aggregations == nil {
		return nil, errors.New("cannot get result aggregation")
	}

	productions, found := result.Aggregations.Composite("production")
	if !found {
		return nil, errors.New("cannot get result composite performance alarm")
	}

	items = append(items, productions.Buckets...)
	return items, nil
}

func (r *solarRepo) GetPerformanceByDate(date *time.Time, efficiencyFactor float64, focusHour int, thresholdPct float64) ([]*elastic.AggregationBucketCompositeItem, error) {
	ctx := context.Background()
	compositeAggregation := elastic.NewCompositeAggregation().
		Size(10000).
		Sources(elastic.NewCompositeAggregationDateHistogramValuesSource("date").Field("@timestamp").CalendarInterval("day").Format("yyyy-MM-dd"),
			elastic.NewCompositeAggregationTermsValuesSource("vendor_type").Field("vendor_type.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("id").Field("id.keyword")).
		SubAggregation("daily", elastic.NewMaxAggregation().Field("daily_production")).
		SubAggregation("capacity", elastic.NewAvgAggregation().Field("installed_capacity")).
		SubAggregation("threshold", elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"capacity": "capacity"}).
			Script(elastic.NewScript("params.capacity * params.efficiency_factor * params.focus_hour * params.threshold_percentage").
				Params(map[string]interface{}{
					"efficiency_factor":    efficiencyFactor,
					"focus_hour":           focusHour,
					"threshold_percentage": thresholdPct,
				}))).
		SubAggregation("hits", elastic.NewTopHitsAggregation().
			Size(1).
			FetchSourceContext(
				elastic.NewFetchSourceContext(true).Include(
					"id", "name", "vendor_type", "node_type", "ac_phase", "plant_status",
					"area", "site_id", "site_city_code", "site_city_name", "installed_capacity",
				)))

	collectorIndex := fmt.Sprintf("%s-%s", r.config.SolarIndex, date.Format("2006.01.02"))
	searchQuery := r.elastic.Search(collectorIndex).
		Size(0).
		Query(elastic.NewBoolQuery().Must(
			elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
			elastic.NewTermsQuery("vendor_type.keyword",
				strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
				strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
				strings.ToUpper(constant.VENDOR_TYPE_INVT),
				strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
			),
		)).
		Aggregation("performance_alarm", compositeAggregation)

	items := make([]*elastic.AggregationBucketCompositeItem, 0)
	result, err := searchQuery.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if result.Aggregations == nil {
		return nil, errors.New("cannot get result aggregations")
	}

	performanceAlarm, found := result.Aggregations.Composite("performance_alarm")
	if !found {
		return nil, errors.New("cannot get result composite performance alarm")
	}

	items = append(items, performanceAlarm.Buckets...)
	if len(performanceAlarm.AfterKey) > 0 {
		afterKey := performanceAlarm.AfterKey

		for {
			query := r.searchIndex().
				Size(0).
				Query(elastic.NewBoolQuery().Must(
					elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
					elastic.NewTermsQuery("vendor_type.keyword",
						strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
						strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
						strings.ToUpper(constant.VENDOR_TYPE_INVT),
						strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
					),
				)).
				Aggregation("performance_alarm", compositeAggregation.AggregateAfter(afterKey))

			result, err := query.Pretty(true).Do(ctx)
			if err != nil {
				return nil, err
			}

			if result.Aggregations == nil {
				return nil, errors.New("cannot get result aggregations")
			}

			performanceAlarm, found := result.Aggregations.Composite("performance_alarm")
			if !found {
				return nil, errors.New("cannot get result composite performance alarm")
			}

			items = append(items, performanceAlarm.Buckets...)

			if len(performanceAlarm.AfterKey) == 0 {
				break
			}

			afterKey = performanceAlarm.AfterKey
		}
	}

	return items, nil

}

func (r *solarRepo) GetPerformanceLow(duration int, efficiencyFactor float64, focusHour int, thresholdPct float64) ([]*elastic.AggregationBucketCompositeItem, error) {
	ctx := context.Background()
	compositeAggregation := elastic.NewCompositeAggregation().
		Size(10000).
		Sources(elastic.NewCompositeAggregationDateHistogramValuesSource("date").Field("@timestamp").CalendarInterval("day").Format("yyyy-MM-dd"),
			elastic.NewCompositeAggregationTermsValuesSource("vendor_type").Field("vendor_type.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("id").Field("id.keyword")).
		SubAggregation("max_daily", elastic.NewMaxAggregation().Field("daily_production")).
		SubAggregation("avg_capacity", elastic.NewAvgAggregation().Field("installed_capacity")).
		SubAggregation("threshold_percentage", elastic.NewBucketScriptAggregation().
			BucketsPathsMap(map[string]string{"capacity": "avg_capacity"}).
			Script(elastic.NewScript("params.capacity * params.efficiency_factor * params.focus_hour * params.threshold_percentage").
				Params(map[string]interface{}{
					"efficiency_factor":    efficiencyFactor,
					"focus_hour":           focusHour,
					"threshold_percentage": thresholdPct,
				}))).
		SubAggregation("under_threshold", elastic.NewBucketSelectorAggregation().
			BucketsPathsMap(map[string]string{"threshold": "threshold_percentage", "daily": "max_daily"}).
			Script(elastic.NewScript("params.daily <= params.threshold"))).
		SubAggregation("hits", elastic.NewTopHitsAggregation().
			Size(1).
			FetchSourceContext(
				elastic.NewFetchSourceContext(true).Include(
					"id", "name", "vendor_type", "node_type", "ac_phase", "plant_status",
					"area", "site_id", "site_city_code", "site_city_name", "installed_capacity",
				)))

	searchQuery := r.searchIndex().
		Size(0).
		Query(elastic.NewBoolQuery().Must(
			elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
			elastic.NewRangeQuery("@timestamp").Gte(fmt.Sprintf("now-%dd/d", duration)).Lte("now-1d/d"),
		)).
		Aggregation("performance_alarm", compositeAggregation)

	items := make([]*elastic.AggregationBucketCompositeItem, 0)
	result, err := searchQuery.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if result.Aggregations == nil {
		return nil, errors.New("cannot get result aggregations")
	}

	performanceAlarm, found := result.Aggregations.Composite("performance_alarm")
	if !found {
		return nil, errors.New("cannot get result composite performance alarm")
	}

	items = append(items, performanceAlarm.Buckets...)
	if len(performanceAlarm.AfterKey) > 0 {
		afterKey := performanceAlarm.AfterKey

		for {
			query := r.searchIndex().
				Size(0).
				Query(elastic.NewBoolQuery().Must(
					elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
					elastic.NewRangeQuery("@timestamp").Gte(fmt.Sprintf("now-%dd/d", duration)).Lte("now-1d/d"),
				)).
				Aggregation("performance_alarm", compositeAggregation.AggregateAfter(afterKey))

			result, err := query.Pretty(true).Do(ctx)
			if err != nil {
				return nil, err
			}

			if result.Aggregations == nil {
				return nil, errors.New("cannot get result aggregations")
			}

			performanceAlarm, found := result.Aggregations.Composite("performance_alarm")
			if !found {
				return nil, errors.New("cannot get result composite performance alarm")
			}

			items = append(items, performanceAlarm.Buckets...)

			if len(performanceAlarm.AfterKey) == 0 {
				break
			}

			afterKey = performanceAlarm.AfterKey
		}
	}

	return items, nil
}

func (r *solarRepo) GetSumPerformanceLow(duration int) ([]*elastic.AggregationBucketCompositeItem, error) {
	ctx := context.Background()
	items := make([]*elastic.AggregationBucketCompositeItem, 0)

	compositeAggregation := elastic.NewCompositeAggregation().
		Size(10000).
		Sources(elastic.NewCompositeAggregationDateHistogramValuesSource("date").Field("@timestamp").CalendarInterval("day").Format("yyyy-MM-dd"),
			elastic.NewCompositeAggregationTermsValuesSource("vendor_type").Field("vendor_type.keyword"),
			elastic.NewCompositeAggregationTermsValuesSource("id").Field("id.keyword")).
		SubAggregation("max_daily", elastic.NewMaxAggregation().Field("daily_production")).
		SubAggregation("avg_capacity", elastic.NewAvgAggregation().Field("installed_capacity")).
		SubAggregation("hits", elastic.NewTopHitsAggregation().
			Size(1).
			FetchSourceContext(
				elastic.NewFetchSourceContext(true).Include(
					"id", "name", "vendor_type", "node_type", "ac_phase", "plant_status",
					"area", "site_id", "site_city_code", "site_city_name", "installed_capacity",
				)))

	searchQuery := r.searchIndex().
		Size(0).
		Query(elastic.NewBoolQuery().Must(
			elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
			elastic.NewRangeQuery("@timestamp").Gte(fmt.Sprintf("now-%dd/d", duration)).Lte("now-1d/d"),
		)).
		Aggregation("performance_alarm", compositeAggregation)

	firstResult, err := searchQuery.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if firstResult.Aggregations == nil {
		return nil, errors.New("cannot get result aggregations")
	}

	performanceAlarm, found := firstResult.Aggregations.Composite("performance_alarm")
	if !found {
		return nil, errors.New("cannot get result composite performance alarm")
	}

	items = append(items, performanceAlarm.Buckets...)

	if len(performanceAlarm.AfterKey) > 0 {
		afterKey := performanceAlarm.AfterKey

		for {
			searchQuery = r.searchIndex().
				Size(0).
				Query(elastic.NewBoolQuery().Must(
					elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
					elastic.NewRangeQuery("@timestamp").Gte(fmt.Sprintf("now-%dd/d", duration)).Lte("now-1d/d"),
				)).
				Aggregation("performance_alarm", compositeAggregation.AggregateAfter(afterKey))

			result, err := searchQuery.Pretty(true).Do(ctx)
			if err != nil {
				return nil, err
			}

			if firstResult.Aggregations == nil {
				return nil, errors.New("cannot get result aggregations")
			}

			performanceAlarm, found := result.Aggregations.Composite("performance_alarm")
			if !found {
				return nil, errors.New("cannot get result composite performance alarm")
			}

			items = append(items, performanceAlarm.Buckets...)

			if len(performanceAlarm.AfterKey) == 0 {
				break
			}

			afterKey = performanceAlarm.AfterKey
		}
	}

	return items, err
}

func (r *solarRepo) GetUniquePlantByIndex(index string) ([]*elastic.AggregationBucketKeyItem, error) {
	ctx := context.Background()
	termAggregation := elastic.NewTermsAggregation().
		Field("name.keyword").
		Size(10000)

	termAggregation = termAggregation.
		SubAggregation(
			"data",
			elastic.
				NewTopHitsAggregation().
				Size(1).
				FetchSourceContext(
					elastic.NewFetchSourceContext(true).
						Include("name", "area", "vendor_type", "installed_capacity", "location", "owner"),
				),
		)

	searchQuery := r.elastic.Search(index).
		Size(0).
		Query(elastic.NewBoolQuery().Must(
			elastic.NewMatchQuery("data_type", constant.DATA_TYPE_PLANT),
		)).
		Aggregation("plant", termAggregation)

	firstResult, err := searchQuery.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if firstResult.Aggregations == nil {
		return nil, errors.New("cannot get result aggregations")
	}

	plant, found := firstResult.Aggregations.Terms("plant")
	if !found {
		return nil, errors.New("cannot get result term plant")
	}

	return plant.Buckets, nil
}

func (r *solarRepo) UpdateOwnerToIndex(index, owner string) error {
	bodyFormat := `{
		"script": {
			"source": "ctx._source.owner = '%v'"
		}
	}`

	body := fmt.Sprintf(bodyFormat, owner)
	_, err := r.elastic.UpdateByQuery(index).Body(body).Do(context.Background())
	return err
}

func (r *solarRepo) DeleteIndex(index string) error {
	deleteIndex, err := r.elastic.DeleteIndex(index).Do(context.Background())
	if err != nil {
		return err
	}

	if !deleteIndex.Acknowledged {
		return fmt.Errorf("[%v]index deletion not acknowledge", index)
	}

	return nil
}

func (r *solarRepo) DeleteDocumentInIndex(index, id string) error {
	_, err := r.elastic.Delete().Index(index).Id(id).Do(context.Background())
	return err
}
