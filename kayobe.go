package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// type User struct {
// 	Id             int64
// 	Email          string `sql:"type:text;"`
// 	HashedPassword string
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// }

// type Site struct {
// 	Domain string `sql:"type:text;"`
// 	Key    string
// }

type Request struct {
	Id               int64
	SessionId        string
	SessionTimestamp int64
	Url              string `sql:"type:text;"`
	Header           string `sql:"type:text;"`
	Domain           string `sql:"type:text;"`
	Referer          string `sql:"type:text;"`
}

const (
	templatesDirectory = "view"
)

var db gorm.DB
var messages chan string

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("view/" + tmpl + ".html")
	t.Execute(w, data)
}

// http://www.golangpatterns.info/web/long-poll-server
func pollHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, <-messages)
	io.WriteString(w, strconv.Itoa(getActiveSessions()))
}

func pushHandler(w http.ResponseWriter, req *http.Request) {

	// body, err := ioutil.ReadAll(req.Body)

	// if err != nil {
	// w.WriteHeader(400)
	// }

	w.WriteHeader(http.StatusOK)
}
func dashboardHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "dashboard", getActiveSessions())
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func getActiveSessions() (count int) {
	var request Request
	db.Last(&request)

	// 15 minute offset
	timeOffset := request.SessionTimestamp - (1000 * 60 * 15)
	var requests []Request
	db.Where("session_timestamp > ?", timeOffset).Find(&requests)
	var metaMap = make(map[string]Request)

	for _, value := range requests {
		metaMap[value.SessionId] = value
	}
	return len(metaMap)
}

func put(str string) {
	fmt.Printf("%s\n", str)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	// return gif
	// base64GifPixel := "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="
	// output, _ := base64.StdEncoding.DecodeString(base64GifPixel)
	output := []byte{71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 255, 255, 255, 0, 0, 0, 44, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 2, 68, 1, 0, 59}
	w.Header().Set("Content-Type", "image/gif")
	io.WriteString(w, string(output))

	go func(req *http.Request) {
		// vars := mux.Vars(req)
		sessionId := req.FormValue("sessionId")
		referer := req.FormValue("referer")
		key := req.FormValue("key")
		// req.FormValue("session")
		if sessionId != "" && referer != "" && key != "" {
			put("Request on " + req.Referer())

			sessionThings := strings.Split(sessionId, "-")
			sessionRandomId := sessionThings[1]
			sessionTimestamp, err := strconv.ParseInt(sessionThings[2], 10, 64)
			if err != nil {
				panic(err)
			}

			request := Request{
				SessionId:        sessionId,
				SessionTimestamp: time.Now().UnixNano(),
				Referer:          referer,
				Url:              req.Referer(),
			}

			query := db.Create(&request)

			if query.Error != nil {
				fmt.Println(query.Error)
			}

			messages <- string("")
		}
	}(req)
}

func main() {
	messages = make(chan string)
	var err error
	db, err = gorm.Open("postgres", "dbname=kayobe sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	query := db.AutoMigrate(&Request{})
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	// accessKey := os.Getenv("KAYOBE_AWS_ACCESS_KEY_ID")
	// secretKey := os.Getenv("KAYOBE_AWS_SECRET_ACCESS_KEY")
	// creds := aws.Creds(accessKey, secretKey, "")
	// dbclient = dynamodb.New(creds, "us-west-2", nil)
	// resp, err := dbclient.ListTables(nil)
	// fmt.Println(resp)

	r := mux.NewRouter()
	r.HandleFunc("/dashboard/", dashboardHandler)
	r.HandleFunc("/push/", pushHandler)
	r.HandleFunc("/poll/", pollHandler)
	r.HandleFunc("/request/", requestHandler)
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Serving site on port :8000")
	http.ListenAndServe(":8000", nil)
}
