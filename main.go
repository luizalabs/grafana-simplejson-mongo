package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/luizalabs/grafana-simplejson-mongo/api"
)

func main() {
	mongoHosts := os.Getenv("MONGO_HOSTS")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	conf := api.Config{
		Port:             port,
		MongoHosts:       strings.Split(mongoHosts, ","),
		CurrentHostIndex: 0,
	}
	errs := make(chan error, 2)
	api.StartHTTPServer(conf, errs)
	log.Println("start")
	for {
		err := <-errs
		log.Println(err)
	}
}
