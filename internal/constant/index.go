package constant

import "strings"

var InitIndex = strings.NewReader(`{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 1,
		"analysis": {
			"analyzer": {
				"my_analyzer": {
					"tokenizer": "my_tokenizer"
				}
			},
			"tokenizer": {
				"my_tokenizer": {
					"type": "ngram",
					"min_gram": 3,
					"max_gram": 3,
					"token_chars": [
						"letter",
						"digit"
					]
				}
			}
		}
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"name": {
				"type": "text",
				"analyzer": "my_analyzer",
				"fields": {
					"keyword": {
						"type": "keyword"
					}
				}
			},
			"description": {
				"type": "text",
				"analyzer": "my_analyzer",
				"fields": {
					"keyword": {
						"type": "keyword",
						"ignore_above": 256
					}
				}
			},
			"price": {
				"type": "double",
				"fields": {
					"keyword": {
						"type": "keyword"
					}
				}
			},
			"owner_id": {
				"type": "text",
				"analyzer": "my_analyzer",
				"fields": {
					"keyword": {
						"type": "keyword"
					}
				}
			},
			"created_at": {
				"type": "date",
				"fields": {
					"keyword": {
						"type": "keyword"
					}
				}
			},
			"updated_at": {
				"type": "date",
				"fields": {
					"keyword": {
						"type": "keyword"
					}
				}
			},
			"deleted_at": {
				"type": "text",
				"fields": {
					"keyword": {
            "null_value": "null",
						"type": "keyword"
					}
				}
			}
		}
	}
}`)

var IndexMapping = strings.NewReader(`{
  "properties": {
    "id": {
      "type": "keyword"
    },
    "name": {
      "type": "text",
      "analyzer": "my_analyzer",
      "fields": {
        "keyword": {
          "type": "keyword"
        }
      }
    },
    "description": {
      "type": "text",
      "analyzer": "my_analyzer",
      "fields": {
        "keyword": {
          "type": "keyword",
          "ignore_above": 256
        }
      }
    },
    "price": {
      "type": "double",
      "fields": {
        "keyword": {
          "type": "keyword"
        }
      }
    },
    "owner_id": {
      "type": "text",
      "analyzer": "my_analyzer",
      "fields": {
        "keyword": {
          "type": "keyword"
        }
      }
    },
    "created_at": {
      "type": "date",
      "fields": {
        "keyword": {
          "type": "keyword"
        }
      }
    },
    "updated_at": {
      "type": "date",
      "fields": {
        "keyword": {
          "type": "keyword"
        }
      }
    },
    "deleted_at": {
      "type": "text",
      "fields": {
        "keyword": {
          "null_value": "null",
          "type": "keyword"
        }
      }
    }
  }
}`)
