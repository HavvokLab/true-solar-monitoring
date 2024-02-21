package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/olivere/elastic/v7"
)

func init() {
	config.InitConfig()
	logger.InitLogger(util.GetLogName("update_altervim_area"))
}

func init() {
	infra.InitGormDB()
}

type AltervimSite struct {
	Area     string `csv:"area"`
	SiteName string `csv:"site_name"`
}

func main() {
	es, err := infra.NewElasticsearch()
	if err != nil {
		logger.GetLogger().Panic(err)
	}

	records := []model.Plant{}
	if err := infra.GormDB.Find(&records).Error; err != nil {
		logger.GetLogger().Panic(err)
	}

	startDate := time.Date(2023, time.October, 1, 0, 0, 0, 0, time.Local)
	dateRange := int(time.Since(startDate).Hours() / 24)
	for i := 0; i < dateRange; i++ {
		date := startDate.AddDate(0, 0, i)
		index := fmt.Sprintf("solarcell-%v", date.Format("2006.01.02"))
		for _, record := range records {
			if record.Area != nil {
				if err := Update(es, index, record.Name, *record.Area); err != nil {
					logger.GetLogger().Error(err)
				}
			}
		}
	}
}

func Update(es *elastic.Client, index string, name string, area string) error {
	bodyFormat := `{
  		"script": {
    		"source": "ctx._source.area = '%v'",
    		"lang": "painless"
  		},
  		"query": {
    		"bool": {
      			"must": [
        			{
          				"term": {
            				"name.keyword": {
              					"value": "%v"
            				}
          				}
        			}
      			]
    		}
  		}
	}`

	body := fmt.Sprintf(bodyFormat, area, name)
	_, err := es.UpdateByQuery(index).Body(body).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
