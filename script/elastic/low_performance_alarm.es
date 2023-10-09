GET solarcell*/_search
{
  "size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "data_type": "PLANT"
          }
        },
        {
          "range": {
            "@timestamp": {
              "gte": "now-7d/d",
              "lte": "now-1d/d"
            }
          }
        }
      ]
    }
  }, 
  "aggs": {
    "performance_alarm": {
      "composite": {
        "size": 10000,
        "sources": [
          {
            "date": {
              "date_histogram": {
                "field": "@timestamp",
                "calendar_interval": "day",
                "format": "yyyy-MM-dd",
                "time_zone": "+07:00"
              }
            }
          },
          {
            "vendor_type": {
              "terms": {
                "field": "vendor_type.keyword"
              }
            }
          },
          {
            "area": {
              "terms": {
                "field": "id.keyword"
              }
            }
          }
        ]
      },
      "aggs": {
        "max_daily": {
          "max": {
            "field": "daily_production"
          }
        },
        "avg_capacity": {
          "max": {
            "field": "installed_capacity"
          }
        },
        "threshold_percentage": {
          "bucket_script": {
            "buckets_path": {
              "capacity": "avg_capacity"
            },
            "script": {
              "params": {
                "efficiency_factor": 0.8,
                "focus_hour": 5,
                "threshold_percentage": 0.6
              },
              "source": "params.capacity * params.efficiency_factor * params.focus_hour * params.threshold_percentage"
            }
          }
        },
        "under_threshold": {
          "bucket_selector": {
            "buckets_path": {
              "threshold": "threshold_percentage",
              "daily": "max_daily"
            },
            "script": "params.daily <= params.threshold"
          }
        },
        "hits":{
          "top_hits": {
            "size": 1,
            "_source": ["id", "name", "vendor_type", "node_type", "ac_phase", "plant_status",
					"area", "site_id", "site_city_code", "site_city_name", "installed_capacity"]
          }
        }
      }
    }
  }
}