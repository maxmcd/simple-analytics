package controllers

import (
	"net/http"

	"github.com/maxmcd/kayobe/util"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	util.RenderTemplate(w, "dashboard", util.GetActiveSessions())
}
