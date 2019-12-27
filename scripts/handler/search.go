package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/verlandz/elasticsearch/scripts/elastic"
	m "github.com/verlandz/elasticsearch/scripts/models"
	"github.com/verlandz/elasticsearch/scripts/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("../views", "search.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Form Value
	q := r.FormValue("q")
	page := r.FormValue("pg")
	ob := r.FormValue("ob")
	fb := r.FormValue("fb")

	// Current Page & Offset
	currPage, _ := strconv.Atoi(page)
	currPage = utils.MaxInt(1, currPage)
	offset := (currPage - 1) * elastic.SIZE

	// Result
	res := elastic.Search(q, offset, ob, fb)
	resStruct := m.ESResult{}
	json.Unmarshal([]byte(res), &resStruct)

	// Min, max, total, hit data
	totalData := resStruct.Hits.Total.Value
	hitData := len(resStruct.Hits.Hits)
	minData := 1 + offset
	maxData := utils.MinInt(totalData, currPage*elastic.SIZE)

	// Max Page
	maxPage := totalData / elastic.SIZE
	if totalData%elastic.SIZE != 0 {
		maxPage++
	}

	var data = map[string]interface{}{
		"result":     res,
		"currSearch": q,
		"currPage":   currPage,
		"maxPage":    maxPage,
		"totalData":  totalData,
		"hitData":    hitData,
		"minData":    minData,
		"maxData":    maxData,
		"currOB":     ob,
		"currFB":     fb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
