package controllers

import "net/http"
import "kayobe/util"

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "index", nil)
}
