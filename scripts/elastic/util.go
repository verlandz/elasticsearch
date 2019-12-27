package elastic

import (
	"fmt"
	"strings"
)

func build_q(q string) string {
	q = strings.ToLower(q)
	q = strings.Join(strings.Split(q, ""), "*")
	if q != "" {
		q = "*" + q + "*"
	}
	return q
}

func build_qSort(ob string) string {
	ob = strings.ToLower(ob)
	qSort := ``
	switch ob {
	case "name":
		qSort += `,{"name.keyword":"asc"}`
	case "date":
		qSort += `,{"release.date":"asc"}`
	case "version":
		qSort += `,{"release.version":"asc"}`
	case "location":
		qSort += `,{"release.location.keyword":"asc"}`
	}
	return qSort
}

func build_qFilter(fb string) string {
	fb = strings.ToLower(fb)
	arr := strings.Split(fb, ",")

	temp := `
	{
		"match": {
			"attribute.name": "%v"
		}
	}
	`
	res, flag := "", false
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" {
			continue
		}
		if flag {
			res += ","
		}
		res += fmt.Sprintf(temp, arr[i])
		flag = true
	}

	return res
}

func build_shouldMatch(qFilter string) string {
	if qFilter != "" {
		return `,"minimum_should_match" : 1`
	}
	return ""
}
