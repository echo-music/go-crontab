package main

import (
	"github.com/echo-music/go-crontab/master/api"
	"log"
	"net/http"
	"time"
)
var  port  = ":8080"
func main() {

	s := &http.Server{
		Addr:           port,
		Handler:        api.RegisterRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Printf("listen on %s", port)
	log.Fatal(s.ListenAndServe())
}
