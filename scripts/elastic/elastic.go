package elastic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	m "github.com/verlandz/elasticsearch/scripts/models"
	"github.com/verlandz/elasticsearch/scripts/utils"
)

// We use :9201 as target
const URL = "http://localhost:9201/robots_v1/"
const SIZE = 30

// Mostly only this type header that used
var header = map[string]string{
	"Content-Type": "application/json",
}

var template = `
	{
		"query": {
			"bool": {
				"must": [
					{
						"wildcard": {
							"name": {
								"value": "%v"
							}
						}
					}
				],
				"should":[
					%v
				]
				%v
			}
		},
		"aggs": {
			"attribute": {
				"terms": {
					"field": "attribute.name.keyword",
					"size":1000
				}
			},
			"ability": {
				"terms": {
					"field": "ability.keyword",
					"size":1000
				}
			}
		},
		"sort" : [
			{"_score":"desc"}
			%v
		],
		"from":%v,
		"size":%v
	}
`

func Insert(i int, robot m.Robot) {
	url := URL + "_doc/" + strconv.Itoa(i)

	marsh, _ := json.Marshal(robot)
	body := strings.NewReader(string(marsh))

	_, err := utils.REST_API(http.MethodPut, url, header, body)
	if err != nil {
		fmt.Printf("[%v] %v\n", i, err)
	} else {
		fmt.Printf("[%v] %v\n", i, robot)
	}
}

func Search(q string, offset int, ob, fb string) string {
	url := URL + "_search"

	q = build_q(q)
	qSort := build_qSort(ob)
	qFilter := build_qFilter(fb)
	shouldMatch := build_shouldMatch(qFilter)

	s := fmt.Sprintf(template,
		q, qFilter, shouldMatch, qSort, offset, SIZE)
	fmt.Println(s)

	body := strings.NewReader(s)
	res, err := utils.REST_API(http.MethodGet, url, header, body)
	if err != nil {
		log.Println(err)
	}

	return string(res)
}
