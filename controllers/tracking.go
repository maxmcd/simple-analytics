package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/maxmcd/kayobe/model"
	"github.com/maxmcd/kayobe/util"
)

var messages chan string

func init() {
	messages = make(chan string)
}

func PollHandler(w http.ResponseWriter, req *http.Request) {
	activeSessions := util.GetActiveSessions()
	activeSessionsString := strconv.Itoa(activeSessions)
	io.WriteString(w, activeSessionsString)
}

func RequestsHandler(w http.ResponseWriter, req *http.Request) {
	var requests []model.Request
	query := model.DB.Find(&requests)
	if query.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonRequestsBytes, err := json.Marshal(requests)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// sentActiveSessions := req.FormValue("activeSessions")

	// activeSessions := util.GetActiveSessions()
	io.WriteString(w, string(jsonRequestsBytes))
}

func TrackingHandler(w http.ResponseWriter, req *http.Request) {

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
			fmt.Println("Request on " + req.Referer())

			headerJson, err := json.Marshal(req.Header)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("requests:", time.Now().UnixNano())
			request := model.Request{
				SessionId:      sessionId,
				Referer:        referer,
				Url:            req.Referer(),
				Ip:             req.RemoteAddr,
				Header:         string(headerJson),
				UserAgent:      req.Header["User-Agent"][0],
				AcceptLanguage: req.Header["Accept-Language"][0],
			}

			query := model.DB.Create(&request)

			if query.Error != nil {
				fmt.Println(query.Error)
			}

			// messages <- string("")
		}
	}(req)
}
