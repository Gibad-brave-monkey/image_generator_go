package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Gibad_brave_monket/image_generator_go/configs"
	"github.com/Gibad_brave_monket/image_generator_go/pkg/img"
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
	buffer, err := img.GenerateFavicon()

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err = w.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server starting in the 8080 PORT")
		if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
			log.Fatal(err)
		}
	}()

	signalValue := <-sigs
	signal.Stop(sigs)

	log.Println("stop signal: ", signalValue)
}
