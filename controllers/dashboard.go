package controllers

import (
	"kayobe/util"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	util.RenderTemplate(w, "dashboard", util.GetActiveSessions())
}
