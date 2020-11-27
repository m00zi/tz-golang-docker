package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func seeTimezone(w http.ResponseWriter, req *http.Request) {
	timeZone := os.Getenv("TIME_ZONE")
	tz, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	localTime := time.Now().In(tz)
	fmt.Fprintf(w, localTime.String())
}

func main() {
	http.HandleFunc("/", seeTimezone)
	http.ListenAndServe(":80", nil)
}
