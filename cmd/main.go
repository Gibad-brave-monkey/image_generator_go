package main

import (
	"flag"
	"log"

	"github.com/Gibad_brave_monket/image_generator_go/configs"
	"github.com/Gibad_brave_monket/image_generator_go/internal/server"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to config .evn")

func main() {
	conf, err := configs.New(*confPath)

	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}
