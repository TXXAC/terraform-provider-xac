terraform {
  required_providers {
    xac = {
      source  = "TXXAC/xac"
      version = "0.1.6"
    }
  }
}

provider "xac123" {
  config_version = "v0.1.0"

  name = "xac123"
}

resource "xac123" "default"{
  deploy = {
    "app" = "study"
    "deploy_env" = "formal"
    "deploy_policy" = "formal_policy"
    "product_type" = "image"
    "product_url" = ""
    "rules" = {}
    "server" = "helloworld"
  }
  deploy_policy = {
    "formal_policy" = {
      "deploy_failure_events" = {
        "event_params" = {
          "content" = "%pipelineURL%发布失败"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "deploy_failure_events" = {
        "event_params" = {
          "confirm_user" = ""
          "timeout" = "6d"
          "timeout_result" = "pass"
        }
        "event_type" = "confirm"
      }
      "deploy_failure_events" = {
        "event_params" = {
          "batch_interval" = 10
          "batch_num" = 0
          "fail_policy" = "continue"
          "image_version" = ""
          "roll_back_type" = "lastVersion"
          "work_wx_receiver" = ""
        }
        "event_type" = "rollBackImage"
      }
      "deploy_start_events" = {
        "event_params" = {
          "content" = "开始%pipelineURL%发布"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "deploy_success_events" = {
        "event_params" = {
          "content" = "%pipelineURL%发布成功"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "stage_failure_events" = {
        "event_params" = {
          "confirm_user" = ""
          "timeout" = "6d"
          "timeout_result" = "pass"
        }
        "event_type" = "confirm"
      }
      "stage_start_events" = []
      "stage_success_events" = {
        "event_params" = {
          "confirm_user" = ""
          "timeout" = "6d"
          "timeout_result" = "pass"
        }
        "event_type" = "confirm"
      }
      "stages" = {
        "check_rules" = {
          "assert_str" = "[err_code] == 0 \u0026\u0026 [data.data_url] != \"\" \u0026\u0026 [data.check_status] == true"
          "check_type" = "Http"
          "desc" = "XAC_007检查"
          "interval_time" = "60s"
          "necessary_observation" = {
            "at_least_observe_time" = ""
            "observation_period" = ""
          }
          "observe_time" = "180s"
          "request" = {
            "body" = {
              "check_begin_time" = "%endTimeUnix007%"
              "cmp_end_time" = "%startTimeUnix007%"
              "info" = {
                "task_id" = "%taskID007%"
              }
              "promql_rule" = {
                "desc" = ""
                "promql_expr" = ""
              }
            }
            "header" = {
              "Content-Type" = "application/json"
              "X-Gateway-Stage" = "RELEASE"
            }
            "method" = "POST"
            "query_map" = {}
            "url" = ""
          }
          "timeout" = "60s"
        }
        "nodes" = {
          "city" = ""
          "node_type" = "canary"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "wait" = "180s"
      }
      "stages" = {
        "check_rules" = {
          "assert_str" = "[err_code] == 0 \u0026\u0026 [data.data_url] != \"\" \u0026\u0026 [data.check_status] == true"
          "check_type" = "Http"
          "desc" = "XAC_007检查"
          "interval_time" = "60s"
          "necessary_observation" = {
            "at_least_observe_time" = ""
            "observation_period" = ""
          }
          "observe_time" = "180s"
          "request" = {
            "body" = {
              "check_begin_time" = "%endTimeUnix007%"
              "cmp_end_time" = "%startTimeUnix007%"
              "info" = {
                "task_id" = "%taskID007%"
              }
              "promql_rule" = {
                "desc" = ""
                "promql_expr" = ""
              }
            }
            "header" = {
              "Content-Type" = "application/json"
              "X-Gateway-Stage" = "RELEASE"
            }
            "method" = "POST"
            "query_map" = {}
            "url" = ""
          }
          "timeout" = "60s"
        }
        "nodes" = {
          "city" = ""
          "node_type" = "canary"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "nodes" = {
          "city" = "sh"
          "node_type" = "normal"
          "num" = 0
          "scale" = "50%"
          "set" = ""
        }
        "nodes" = {
          "city" = "sz"
          "node_type" = "normal"
          "num" = 0
          "scale" = "50%"
          "set" = ""
        }
        "wait" = "180s"
      }
      "stages" = {
        "check_rules" = {
          "assert_str" = "[err_code] == 0 \u0026\u0026 [data.data_url] != \"\" \u0026\u0026 [data.check_status] == true"
          "check_type" = "Http"
          "desc" = "XAC_007检查"
          "interval_time" = "60s"
          "necessary_observation" = {
            "at_least_observe_time" = ""
            "observation_period" = ""
          }
          "observe_time" = "180s"
          "request" = {
            "body" = {
              "check_begin_time" = "%endTimeUnix007%"
              "cmp_end_time" = "%startTimeUnix007%"
              "info" = {
                "task_id" = "%taskID007%"
              }
              "promql_rule" = {
                "desc" = ""
                "promql_expr" = ""
              }
            }
            "header" = {
              "Content-Type" = "application/json"
              "X-Gateway-Stage" = "RELEASE"
            }
            "method" = "POST"
            "query_map" = {}
            "url" = ""
          }
          "timeout" = "60s"
        }
        "nodes" = {
          "city" = ""
          "node_type" = "canary"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "nodes" = {
          "city" = "sh"
          "node_type" = "normal"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "nodes" = {
          "city" = "sz"
          "node_type" = "normal"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "wait" = "180s"
      }
      "unavailable_node_max_ratio" = "20%"
    }
    "sandbox_policy" = {
      "deploy_failure_events" = {
        "event_params" = {
          "content" = "%pipelineURL%发布失败"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "deploy_failure_events" = {
        "event_params" = {
          "confirm_user" = ""
        }
        "event_type" = "confirm"
      }
      "deploy_start_events" = {
        "event_params" = {
          "content" = "开始%pipelineURL%发布"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "deploy_success_events" = {
        "event_params" = {
          "content" = "%pipelineURL%发布成功"
          "receiver" = ""
        }
        "event_type" = "sendWorkWx"
      }
      "stage_failure_events" = {
        "event_params" = {
          "confirm_user" = ""
        }
        "event_type" = "confirm"
      }
      "stage_start_events" = []
      "stage_success_events" = []
      "stages" = {
        "check_rules" = {
          "assert_str" = "[err_code] == 0 \u0026\u0026 [data.data_url] != \"\" \u0026\u0026 [data.check_status] == true"
          "check_type" = "Http"
          "desc" = "XAC_007检查"
          "interval_time" = "60s"
          "necessary_observation" = {
            "at_least_observe_time" = ""
            "observation_period" = ""
          }
          "observe_time" = "180s"
          "request" = {
            "body" = {
              "check_begin_time" = "%endTimeUnix007%"
              "cmp_end_time" = "%startTimeUnix007%"
              "info" = {
                "task_id" = "%taskID007%"
              }
              "promql_rule" = {
                "desc" = ""
                "promql_expr" = ""
              }
            }
            "header" = {
              "Content-Type" = "application/json"
              "X-Gateway-Stage" = "RELEASE"
            }
            "method" = "POST"
            "query_map" = {}
            "url" = ""
          }
          "timeout" = "60s"
        }
        "nodes" = {
          "city" = ""
          "node_type" = "normal"
          "num" = 0
          "scale" = "100%"
          "set" = ""
        }
        "wait" = ""
      }
      "unavailable_node_max_ratio" = "100%"
    }
  }
}
