package server

import (
	"log"
	"net/http"

	"github.com/Gibad_brave_monket/image_generator_go/configs"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "img")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.icon", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	log.Println("Server starting in the 8080 PORT")

	if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
		log.Fatal(err)
	}
}
