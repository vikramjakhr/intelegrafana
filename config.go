package main

const DashboardJson = `{
    "dashboard":{
        "annotations":{
            "list":[
                {
                    "builtIn":1,
                    "datasource":"-- Grafana --",
                    "enable":true,
                    "hide":true,
                    "iconColor":"rgba(0, 211, 255, 1)",
                    "name":"Annotations & Alerts",
                    "type":"dashboard"
                }
            ]
        },
        "description":"(\"server single\")",
        "editable":true,
        "gnetId":1375,
        "graphTooltip":0,
        "hideControls":false,
        "id":null,
        "links":[
            {
                "asDropdown":true,
                "icon":"external link",
                "includeVars":false,
                "tags":[

                ],
                "title":"Dashboards",
                "type":"dashboards"
            }
        ],
        "refresh":"1m",
        "rows":[

        ],
        "schemaVersion":14,
        "style":"dark",
        "tags":[

        ],
        "templating":{
            "list":[
                {
                    "allValue":null,
                    "current":{

                    },
                    "datasource":"$DATASOURCE_NAME$",
                    "hide":0,
                    "includeAll":false,
                    "label":null,
                    "multi":false,
                    "name":"host",
                    "options":[

                    ],
                    "query":"SHOW TAG VALUES FROM system WITH KEY = \"host\"",
                    "refresh":1,
                    "regex":"",
                    "sort":1,
                    "tagValuesQuery":null,
                    "tags":[

                    ],
                    "tagsQuery":null,
                    "type":"query",
                    "useTags":false
                }
            ]
        },
        "time":{
            "from":"now-5m",
            "to":"now"
        },
        "timepicker":{
            "refresh_intervals":[
                "5s",
                "10s",
                "30s",
                "1m",
                "5m",
                "15m",
                "30m",
                "1h",
                "2h",
                "1d"
            ],
            "time_options":[
                "5m",
                "15m",
                "1h",
                "6h",
                "12h",
                "24h",
                "2d",
                "7d",
                "30d"
            ]
        },
        "timezone":"browser",
        "title":"$DASHBOARD_NAME$",
        "version":0
    },
    "overwrite":false
}`

const PortPanel = `{
    "cacheTimeout":null,
    "colorBackground":true,
    "colorValue":false,
    "colors":[
        "#299c46",
        "rgba(237, 129, 40, 0.89)",
        "#d44a3a"
    ],
    "datasource":"$DATASOURCE_NAME$",
    "format":"none",
    "gauge":{
        "maxValue":100,
        "minValue":0,
        "show":false,
        "thresholdLabels":false,
        "thresholdMarkers":true
    },
    "interval":null,
    "links":[

    ],
    "mappingType":1,
    "mappingTypes":[
        {
            "name":"value to text",
            "value":1
        },
        {
            "name":"range to text",
            "value":2
        }
    ],
    "maxDataPoints":100,
    "nullPointMode":"connected",
    "nullText":null,
    "postfix":"",
    "postfixFontSize":"50%",
    "prefix":"Port ($PORT$): ",
    "prefixFontSize":"30%",
    "rangeMaps":[
        {
            "from":"null",
            "text":"N/A",
            "to":"null"
        }
    ],
    "span":4,
    "sparkline":{
        "fillColor":"rgba(31, 118, 189, 0.18)",
        "full":false,
        "lineColor":"rgb(31, 120, 193)",
        "show":false
    },
    "tableColumn":"",
    "targets":[
        {
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"net_response",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "result_type"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"last"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                },
                {
                    "condition":"AND",
                    "key":"port",
                    "operator":"=",
                    "value":"$PORT$"
                }
            ]
        }
    ],
    "thresholds":"1,1",
    "title":"Port ($PORT$)",
    "type":"singlestat",
    "valueFontSize":"30%",
    "valueMaps":[
        {
            "op":"=",
            "text":"OK",
            "value":"0"
        },
        {
            "op":"=",
            "text":"Down",
            "value":"1"
        }
    ],
    "valueName":"current"
}`

const DiskPanel = `{
    "aliasColors":{
        "undefined":"#1F78C1"
    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "decimals":null,
    "editable":true,
    "error":false,
    "fill":0,
    "legend":{
        "avg":false,
        "current":false,
        "max":false,
        "min":false,
        "show":true,
        "total":false,
        "values":false
    },
    "lines":true,
    "linewidth":2,
    "links":[

    ],
    "minSpan":2,
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":3,
    "points":false,
    "renderer":"flot",
    "repeat":null,
    "seriesOverrides":[
        {
            "alias":"/^inodes.*/i",
            "color":"#82B5D8",
            "yaxis":2
        },
        {
            "alias":"/^space.*/i",
            "color":"#F4D598"
        },
        {
            "alias":"/^space: Total.*/i",
            "fillBelowTo":"space: free, /",
            "lines":false
        },
        {
            "alias":"/^inodes: total.*/i",
            "fillBelowTo":"inodes: free, /",
            "lines":false
        },
        {
            "alias":"/^space: free.*/i",
            "fill":8,
            "linewidth":4
        },
        {
            "alias":"/^inodes: free.*/i",
            "linewidth":4
        }
    ],
    "spaceLength":10,
    "span":4,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"space: Total, $tag_path",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "path"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"disk",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "total"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"space: free, $tag_path",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "path"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"disk",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"B",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "free"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"inodes: total, $tag_path",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "path"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"disk",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "inodes_total"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"inodes: free, $tag_path",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "path"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"disk",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"D",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "inodes_free"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"DISK partitions",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"bytes",
            "label":"space",
            "logBase":1,
            "max":null,
            "min":"0",
            "show":true
        },
        {
            "format":"short",
            "label":"inodes",
            "logBase":1,
            "max":null,
            "min":"0",
            "show":true
        }
    ]
}`

const ProcessPanel = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "decimals":0,
    "editable":true,
    "error":false,
    "fill":0,
    "legend":{
        "avg":false,
        "current":false,
        "max":false,
        "min":false,
        "show":true,
        "total":false,
        "values":false
    },
    "lines":true,
    "linewidth":2,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":3,
    "points":false,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"running",
            "zindex":3
        }
    ],
    "spaceLength":10,
    "span":4,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"Total",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "total"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"running",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"B",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "running"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"blocked",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "blocked"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"idle",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"D",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "idle"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"zombies",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"E",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "zombies"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"sleeping",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"F",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "sleeping"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"total threads",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"G",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "total_threads"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"unknown",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"H",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "unknown"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"paging",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"processes",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"I",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "paging"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"Processes",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":null,
            "show":true
        },
        {
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":null,
            "show":false
        }
    ]
}`

const Swap = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "editable":true,
    "error":false,
    "fill":0,
    "interval":">5m",
    "legend":{
        "alignAsTable":false,
        "avg":false,
        "current":false,
        "hideEmpty":true,
        "hideZero":true,
        "max":false,
        "min":false,
        "rightSide":false,
        "show":true,
        "total":false,
        "values":false
    },
    "lines":true,
    "linewidth":2,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":5,
    "points":false,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"used",
            "fill":4,
            "linewidth":0,
            "yaxis":2
        }
    ],
    "spaceLength":10,
    "span":4,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"in",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "$interval"
                    ],
                    "type":"time"
                },
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                },
                {
                    "params":[
                        "null"
                    ],
                    "type":"fill"
                }
            ],
            "measurement":"swap",
            "policy":"default",
            "refId":"B",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "in"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"mean"
                    },
                    {
                        "params":[
                            "1s"
                        ],
                        "type":"derivative"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"out",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "$interval"
                    ],
                    "type":"time"
                },
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                },
                {
                    "params":[
                        "null"
                    ],
                    "type":"fill"
                }
            ],
            "measurement":"swap",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "out"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"mean"
                    },
                    {
                        "params":[
                            "1s"
                        ],
                        "type":"derivative"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"used",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "$interval"
                    ],
                    "type":"time"
                },
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                },
                {
                    "params":[
                        "null"
                    ],
                    "type":"fill"
                }
            ],
            "measurement":"swap",
            "policy":"default",
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "used_percent"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"mean"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[
        {
            "colorMode":"critical",
            "fill":true,
            "line":true,
            "op":"gt"
        }
    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"swap",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"Bps",
            "label":"I/O",
            "logBase":1,
            "max":null,
            "min":"0",
            "show":true
        },
        {
            "format":"percent",
            "label":"used",
            "logBase":1,
            "max":"100",
            "min":"0",
            "show":true
        }
    ]
}`

const CPUPanel = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "editable":true,
    "error":false,
    "fill":0,
    "hideTimeOverride":false,
    "interval":">2s",
    "legend":{
        "alignAsTable":true,
        "avg":true,
        "current":true,
        "max":true,
        "min":false,
        "rightSide":true,
        "show":true,
        "total":false,
        "values":true
    },
    "lines":true,
    "linewidth":1,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":5,
    "points":false,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"/cores/",
            "linewidth":2,
            "stack":false,
            "yaxis":2,
            "zindex":-2
        },
        {
            "alias":"IDLE",
            "color":"#629E51",
            "stack":false
        },
        {
            "alias":"usage_iowait",
            "color":"#890F02",
            "fill":6,
            "fillBelowTo":"usage_user"
        },
        {
            "alias":"usage_user",
            "fillBelowTo":"usage_system"
        },
        {
            "alias":"usage_system",
            "fill":5
        }
    ],
    "spaceLength":10,
    "span":8,
    "stack":true,
    "steppedLine":false,
    "targets":[
        {
            "alias":"usage_system",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "hide":false,
            "measurement":"cpu",
            "orderByTime":"ASC",
            "policy":"default",
            "query":"SELECT derivative(mean(\"time_iowait\"), 2s) FROM \"cpu\" WHERE $timeFilter GROUP BY time($interval) ,\"host\" fill(none)",
            "rawQuery":false,
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "usage_system"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"usage_user",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "hide":false,
            "measurement":"cpu",
            "orderByTime":"ASC",
            "policy":"default",
            "query":"SELECT derivative(mean(\"time_iowait\"), 2s) FROM \"cpu\" WHERE $timeFilter GROUP BY time($interval) ,\"host\" fill(none)",
            "rawQuery":false,
            "refId":"D",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "usage_user"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"usage_iowait",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "hide":false,
            "measurement":"cpu",
            "orderByTime":"ASC",
            "policy":"default",
            "query":"SELECT derivative(mean(\"time_iowait\"), 2s) FROM \"cpu\" WHERE $timeFilter GROUP BY time($interval) ,\"host\" fill(none)",
            "rawQuery":false,
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "usage_iowait"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"CPU usage",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"percent",
            "label":"usage",
            "logBase":1,
            "max":"100",
            "min":"0",
            "show":true
        },
        {
            "format":"short",
            "label":"",
            "logBase":1,
            "max":"10",
            "min":"0",
            "show":true
        }
    ]
}`

const RAMPanel = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "editable":true,
    "error":false,
    "fill":0,
    "legend":{
        "alignAsTable":true,
        "avg":true,
        "current":true,
        "max":true,
        "min":false,
        "show":true,
        "total":false,
        "values":true
    },
    "lines":true,
    "linewidth":2,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":3,
    "points":false,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"Total",
            "fillBelowTo":"Available",
            "lines":false
        },
        {
            "alias":"Available",
            "color":"#7EB26D",
            "fill":2
        }
    ],
    "spaceLength":10,
    "span":4,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"Available",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"mem",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "available"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"Total",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"mem",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"B",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "total"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"RAM",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"bytes",
            "label":null,
            "logBase":1,
            "max":null,
            "min":"0",
            "show":true
        },
        {
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":null,
            "show":true
        }
    ]
}`

const IPTrafficPanel = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "editable":true,
    "error":false,
    "fill":0,
    "interval":">1s",
    "legend":{
        "alignAsTable":true,
        "avg":true,
        "current":false,
        "max":true,
        "min":false,
        "rightSide":true,
        "show":true,
        "total":false,
        "values":true
    },
    "lines":false,
    "linewidth":0,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":1,
    "points":true,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"/[ms]/",
            "fill":1,
            "lines":true,
            "linewidth":2,
            "points":false
        },
        {
            "alias":"/In/i",
            "color":"#1F78C1"
        },
        {
            "alias":"/Out/i",
            "color":"#CCA300"
        }
    ],
    "spaceLength":10,
    "span":6,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"Out (5m)",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "5m"
                    ],
                    "type":"time"
                },
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                },
                {
                    "params":[
                        "null"
                    ],
                    "type":"fill"
                }
            ],
            "measurement":"nstat",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "IpExtOutOctets"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"mean"
                    },
                    {
                        "params":[
                            "1s"
                        ],
                        "type":"derivative"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"In (5m)",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "5m"
                    ],
                    "type":"time"
                },
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                },
                {
                    "params":[
                        "null"
                    ],
                    "type":"fill"
                }
            ],
            "hide":false,
            "measurement":"nstat",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"D",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "IpExtInOctets"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"mean"
                    },
                    {
                        "params":[
                            "1s"
                        ],
                        "type":"derivative"
                    },
                    {
                        "params":[
                            "*(-1)"
                        ],
                        "type":"math"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"IP traffic",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "format":"Bps",
            "label":"",
            "logBase":1,
            "max":null,
            "min":null,
            "show":true
        },
        {
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":null,
            "show":true
        }
    ]
}`

const SystemLoadPanel = `{
    "aliasColors":{

    },
    "bars":false,
    "dashLength":10,
    "dashes":false,
    "datasource":"$DATASOURCE_NAME$",
    "editable":true,
    "error":false,
    "fill":3,
    "legend":{
        "alignAsTable":true,
        "avg":true,
        "current":true,
        "max":true,
        "min":false,
        "rightSide":false,
        "show":true,
        "total":false,
        "values":true
    },
    "lines":true,
    "linewidth":1,
    "links":[

    ],
    "nullPointMode":"connected",
    "percentage":false,
    "pointradius":5,
    "points":false,
    "renderer":"flot",
    "seriesOverrides":[
        {
            "alias":"/.*cores$/",
            "fill":0
        }
    ],
    "spaceLength":10,
    "span":6,
    "stack":false,
    "steppedLine":false,
    "targets":[
        {
            "alias":"load (1m)",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"system",
            "orderByTime":"ASC",
            "policy":"default",
            "query":"SELECT \"load1\" FROM \"system\" WHERE (\"host\" =~ /^$host$/) AND $timeFilter GROUP BY \"host\"",
            "rawQuery":false,
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "load1"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"load (5m)",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"system",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"B",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "load5"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        },
        {
            "alias":"load (15m)",
            "dsType":"influxdb",
            "groupBy":[
                {
                    "params":[
                        "host"
                    ],
                    "type":"tag"
                }
            ],
            "measurement":"system",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"C",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "load15"
                        ],
                        "type":"field"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                }
            ]
        }
    ],
    "thresholds":[

    ],
    "timeFrom":null,
    "timeShift":null,
    "title":"system: load (5m)",
    "tooltip":{
        "msResolution":false,
        "shared":true,
        "sort":0,
        "value_type":"individual"
    },
    "type":"graph",
    "xaxis":{
        "buckets":null,
        "mode":"time",
        "name":null,
        "show":true,
        "values":[

        ]
    },
    "yaxes":[
        {
            "decimals":2,
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":"0",
            "show":true
        },
        {
            "format":"short",
            "label":null,
            "logBase":1,
            "max":null,
            "min":null,
            "show":false
        }
    ]
}`

const URLPanel = `{
  "cacheTimeout": null,
  "colorBackground": true,
  "colorValue": false,
  "colors": [
    "#299c46",
    "rgba(237, 129, 40, 0.89)",
    "#d44a3a"
  ],
  "datasource": "$DATASOURCE_NAME$",
  "format": "none",
  "gauge": {
    "maxValue": 100,
    "minValue": 0,
    "show": false,
    "thresholdLabels": false,
    "thresholdMarkers": true
  },
  "interval": null,
  "links": [],
  "mappingType": 2,
  "mappingTypes": [
    {
      "name": "value to text",
      "value": 1
    },
    {
      "name": "range to text",
      "value": 2
    }
  ],
  "maxDataPoints": 100,
  "nullPointMode": "connected",
  "nullText": null,
  "postfix": "",
  "postfixFontSize": "50%",
  "prefix": "",
  "prefixFontSize": "50%",
  "rangeMaps": [
    {
      "from": "200",
      "text": "OK",
      "to": "200"
    },
    {
      "from": "201",
      "text": "Error",
      "to": "600"
    }
  ],
  "span": 3,
  "sparkline": {
    "fillColor": "rgba(31, 118, 189, 0.18)",
    "full": false,
    "lineColor": "rgb(31, 120, 193)",
    "show": true
  },
  "tableColumn": "",
  "targets": [
    {
      "dsType": "influxdb",
      "groupBy": [],
      "measurement": "http_response",
      "orderByTime": "ASC",
      "policy": "default",
      "refId": "A",
      "resultFormat": "time_series",
      "select": [
        [
          {
            "params": [
              "http_response_code"
            ],
            "type": "field"
          }
        ]
      ],
      "tags": [
        {
          "key": "host",
          "operator": "=~",
          "value": "/^$host$/"
        },
        {
          "condition": "AND",
          "key": "server",
          "operator": "=",
          "value": "$URL$"
        }
      ]
    }
  ],
  "thresholds": "201,302",
  "title": "$URL$",
  "transparent": true,
  "type": "singlestat",
  "valueFontSize": "150%",
  "valueMaps": [
    {
      "op": "=",
      "text": "N/A",
      "value": "null"
    }
  ],
  "valueName": "current"
}`

const PageLoadPanel = `{
  "cacheTimeout": null,
  "colorBackground": false,
  "colorValue": true,
  "colors": [
    "#299c46",
    "rgba(237, 129, 40, 0.89)",
    "#d44a3a"
  ],
  "datasource": "$DATASOURCE_NAME$",
  "format": "s",
  "gauge": {
    "maxValue": 6,
    "minValue": 0,
    "show": true,
    "thresholdLabels": false,
    "thresholdMarkers": true
  },
  "interval": null,
  "links": [],
  "mappingType": 1,
  "mappingTypes": [
    {
      "name": "value to text",
      "value": 1
    },
    {
      "name": "range to text",
      "value": 2
    }
  ],
  "maxDataPoints": 100,
  "nullPointMode": "connected",
  "nullText": null,
  "postfix": "",
  "postfixFontSize": "50%",
  "prefix": "",
  "prefixFontSize": "30%",
  "rangeMaps": [
    {
      "from": "null",
      "text": "N/A",
      "to": "null"
    }
  ],
  "span": 3,
  "sparkline": {
    "fillColor": "rgba(31, 118, 189, 0.18)",
    "full": false,
    "lineColor": "rgb(31, 120, 193)",
    "show": true
  },
  "tableColumn": "",
  "targets": [
    {
      "dsType": "influxdb",
      "groupBy": [],
      "measurement": "http_response",
      "orderByTime": "ASC",
      "policy": "default",
      "refId": "A",
      "resultFormat": "time_series",
      "select": [
        [
          {
            "params": [
              "response_time"
            ],
            "type": "field"
          }
        ]
      ],
      "tags": [
        {
          "key": "host",
          "operator": "=~",
          "value": "/^$host$/"
        },
        {
          "condition": "AND",
          "key": "server",
          "operator": "=",
          "value": "$URL$"
        }
      ]
    }
  ],
  "thresholds": "3,4",
  "title": "Page Load Time ($URL$)",
  "transparent": true,
  "type": "singlestat",
  "valueFontSize": "70%",
  "valueMaps": [
    {
      "op": "=",
      "text": "N/A",
      "value": "null"
    }
  ],
  "valueName": "current"
}`

const Procstat = `{
    "cacheTimeout":null,
    "colorBackground":true,
    "colorValue":false,
    "colors":[
        "#d44a3a",
        "rgba(237, 129, 40, 0.89)",
        "#299c46"
    ],
    "datasource":"$DATASOURCE_NAME$",
    "format":"none",
    "gauge":{
        "maxValue":100,
        "minValue":0,
        "show":false,
        "thresholdLabels":false,
        "thresholdMarkers":true
    },
    "interval":null,
    "links":[

    ],
    "mappingType":2,
    "mappingTypes":[
        {
            "name":"value to text",
            "value":1
        },
        {
            "name":"range to text",
            "value":2
        }
    ],
    "maxDataPoints":100,
    "nullPointMode":"connected",
    "nullText":null,
    "postfix":"",
    "postfixFontSize":"50%",
    "prefix":"Process ($PROCESS$):",
    "prefixFontSize":"30%",
    "rangeMaps":[
        {
            "from":"2",
            "text":"Running",
            "to":"200000"
        }
    ],
    "span":3,
    "sparkline":{
        "fillColor":"rgba(31, 118, 189, 0.18)",
        "full":false,
        "lineColor":"rgb(31, 120, 193)",
        "show":false
    },
    "tableColumn":"",
    "targets":[
        {
            "alias":"",
            "dsType":"influxdb",
            "groupBy":[

            ],
            "measurement":"procstat",
            "orderByTime":"ASC",
            "policy":"default",
            "refId":"A",
            "resultFormat":"time_series",
            "select":[
                [
                    {
                        "params":[
                            "pid"
                        ],
                        "type":"field"
                    },
                    {
                        "params":[

                        ],
                        "type":"last"
                    }
                ]
            ],
            "tags":[
                {
                    "key":"host",
                    "operator":"=~",
                    "value":"/^$host$/"
                },
                {
                    "condition":"AND",
                    "key":"pattern",
                    "operator":"=",
                    "value":"$PROCESS$"
                }
            ]
        }
    ],
    "thresholds":"0,1",
    "title":"Process ($PROCESS$)",
    "type":"singlestat",
    "valueFontSize":"30%",
    "valueMaps":[
        {
            "op":"=",
            "text":"0",
            "value":"null"
        }
    ],
    "valueName":"current"
}`
