
GET solarcell-2023.*/_search
{
    "size": 0,
    "query": {
        "bool": {
            "must": [
                {
                    "term": {
                        "data_type.keyword": "PLANT"
                    }
                },
                {
                    "terms": {
                        "vendor_type.keyword": [
                            "HUAWEI",
                            "INVT",
                            "GROWATT",
                            "KSTAR"
                        ]
                    }
                },
                {
                    "range": {
                        "@timestamp": {
                            "gte": "now-365d/d",
                            "lte": "now-1d/d"
                        }
                    }
                }
            ]
        }
    },
    "aggs": {
        "data": {
            "composite": {
                "size": 10,
                "sources": [
                    {
                        "date": {
                            "date_histogram": {
                                "field": "@timestamp",
                                "calendar_interval": "month",
                                "format": "yyyy-MM-dd",
                                "time_zone": "Asia/Bangkok"
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
                                "field": "area.keyword"
                            }
                        }
                    },
                    {
                        "name": {
                            "terms": {
                                "field": "name.keyword"
                            }
                        }
                    }
                ]
            },
            "aggs": {
                "hits": {
                    "top_hits": {
                        "size": 1,
                        "_source": [
                            "lat",
                            "lng"
                        ]
                    }
                },
                "installed_capacity": {
                    "max": {
                        "field": "installed_capacity"
                    }
                },
                "monthly_production": {
                    "max": {
                        "field": "monthly_production"
                    }
                },
                "daily_production": {
                    "max": {
                        "field": "daily_production"
                    }
                },
                "target": {
                    "bucket_script": {
                        "buckets_path": {
                            "installed_capacity": "installed_capacity",
                            "daily_production": "daily_production"
                        },
                        "script": "if (params.installed_capacity == 0 || params.daily_production == 0 ) { return 0 } else { (params.installed_capacity*5*0.8*31) }"
                    }
                },
                "production_to_target": {
                    "bucket_script": {
                        "buckets_path": {
                            "installed_capacity": "installed_capacity",
                            "monthly_production": "monthly_production"
                        },
                        "script": "if (params.installed_capacity == 0 || params.monthly_production == 0 ) { return 0 } else { (params.monthly_production/(params.installed_capacity*5*0.8*31))*100 }"
                    }
                }
            }
        }
    }
}