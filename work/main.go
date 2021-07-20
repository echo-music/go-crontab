package main

import (
	"github.com/echo-music/go-crontab/work/logic"
	"log"
	"net/http"
	"time"
)
var  port  = ":8081"
func main() {

	logic.Watcher()
	s := &http.Server{
		Addr:           port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Printf("listen on %s", port)
	log.Fatal(s.ListenAndServe())
}
