package controllers

import (
	"net/http"

	"github.com/maxmcd/kayobe/util"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "index", nil)
}
