package util

import (
	"net/http"
	"text/template"
	"time"

	"github.com/maxmcd/kayobe/model"
)

var T = template.Must(template.ParseGlob("view/*"))

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	T.ExecuteTemplate(w, tmpl+".html", data)
}

func GetActiveSessions() (count int) {

	// 15 minute offset
	timeOffset := time.Now().Add(-time.Duration(time.Minute * 10))

	var requests []model.Request
	model.DB.Where("created_at > ?", timeOffset).Find(&requests)
	var metaMap = make(map[string]model.Request)

	for _, value := range requests {
		metaMap[value.SessionId] = value
	}
	return len(metaMap)
}
