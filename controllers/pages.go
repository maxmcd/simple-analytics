package controllers

import (
	"kayobe/util"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "index", nil)
}
