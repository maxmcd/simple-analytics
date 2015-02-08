package controllers

import (
	"net/http"

	"github.com/maxmcd/kayobe/model"
	"github.com/maxmcd/kayobe/util"
)

type AuthData struct {
	Name   string
	Submit string
	Action string
	Email  string
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	data := AuthData{
		Name:   "Log In",
		Submit: "Log In",
		Action: "login",
	}

	switch req.Method {

	case "GET":
		util.RenderTemplate(w, "login_signup", data)

	case "POST":
		// if post
		email := req.FormValue("email")
		password := req.FormValue("password")

		user, err := model.ValidateUserPassword(email, password)
		_ = user
		if err != nil {
			// whoops
			data.Email = email
			util.RenderTemplate(w, "login_signup", data)
		} else {
			http.Redirect(w, req, "/dashboard/", 302)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func SignupHandler(w http.ResponseWriter, req *http.Request) {
	data := AuthData{
		Name:   "Sign Up",
		Submit: "Sign Up",
		Action: "sign-up",
	}

	switch req.Method {

	case "GET":
		util.RenderTemplate(w, "login_signup", data)

	case "POST":
		// if post
		email := req.FormValue("email")
		password := req.FormValue("password")

		if email == "" || password == "" {
			// whoops
			data.Email = email
			util.RenderTemplate(w, "login_signup", data)
			return
		}

		user, err := model.NewUser(email, password)
		_ = user

		if err != nil {
			data.Email = email
			util.RenderTemplate(w, "login_signup", data)
			return
		}

		http.Redirect(w, req, "/dashboard/", 302)
	default:
		w.WriteHeader(http.StatusNotFound)

	}
}
