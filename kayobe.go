package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxmcd/kayobe/controllers"
)

const (
	templatesDirectory = "view"
)

// http://www.golangpatterns.info/web/long-poll-server

func puts(str string) {
	fmt.Printf("%s\n", str)
}

func main() {

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/dashboard/", controllers.DashboardHandler)
	r.HandleFunc("/poll/", controllers.PollHandler)
	r.HandleFunc("/api/{command}/", controllers.ApiHandler)
	r.HandleFunc("/request/", controllers.TrackingHandler)
	r.HandleFunc("/requests/", controllers.RequestsHandler)
	r.HandleFunc("/login/", controllers.LoginHandler)
	r.HandleFunc("/sign-up/", controllers.SignupHandler)
	r.HandleFunc("/", controllers.IndexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)

	fmt.Println("Serving site on port :8000")
	http.ListenAndServe(":8000", nil)
}
