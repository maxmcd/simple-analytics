package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/dynamodb"
	"github.com/gorilla/mux"
)

const (
	templatesDirectory = "view"
)

var templates *template.Template

func displayTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl+".html", data)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	displayTemplate(w, "dashboard", nil)
}

func initializeTemplates() {
	var templateFiles string
	files, err := ioutil.ReadDir(templatesDirectory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		templateFiles = templatesDirectory + "/" + file.Name()
	}
	templates = template.Must(template.ParseFiles(templateFiles))
}

// http://www.golangpatterns.info/web/long-poll-server
func pollHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, <-messages)
}

func pushHandler(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(400)
	}
	messages <- string(body)
	w.WriteHeader(http.StatusOK)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println(key)
	}()
	base64GifPixel := "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="
	w.Header().Set("Content-Type", "image/gif")
	output, _ := base64.StdEncoding.DecodeString(base64GifPixel)
	io.WriteString(w, string(output))
}

var messages chan string
var dbclient *dynamodb.DynamoDB

func main() {
	messages = make(chan string)

	initializeTemplates()

	accessKey := os.Getenv("KAYOBE_AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("KAYOBE_AWS_SECRET_ACCESS_KEY")
	creds := aws.Creds(accessKey, secretKey, "")
	dbclient = dynamodb.New(creds, "us-west-2", nil)
	resp, err := dbclient.ListTables(nil)

	fmt.Println(resp)

	r := mux.NewRouter()
	r.HandleFunc("/dashboard", dashboardHandler)
	r.HandleFunc("/push", pushHandler)
	r.HandleFunc("/poll", pollHandler)
	r.HandleFunc("/request/{key}/", requestHandler)
	http.Handle("/", r)
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Serving site on port :8080")
	http.ListenAndServe(":8080", nil)
}
