package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var thresholdTime time.Time
var rickRollUrl = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
var secretUrl string

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	// Compare time.Now() with the time in the environment variable THRESHOLD_TIME
	// If the time is greater than the threshold time, return a 500 status code
	// Otherwise, return a 200 status code
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Set the content type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	pastThreshold := time.Now().After(thresholdTime)
	if pastThreshold {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("{\"url\":\"" + secretUrl + "\",\"valid\":true}"))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		_, err := w.Write([]byte("{\"url\":\"" + rickRollUrl + "\",\"valid\":false}"))
		if err != nil {
			return
		}
	}
}

func loadEnv() {
	// Get the environment variable THRESHOLD_TIME
	// Convert the value to a time.Time object
	// Assign the value to the global variable thresholdTime
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "THRESHOLD_TIME" {
			result, err := time.Parse("2006-01-02T15:04:05-0700", pair[1])
			if err != nil {
				log.Fatal(err)
			}
			thresholdTime = result

		}
		if pair[0] == "SECRET_URL" {
			secretUrl = pair[1]
		}
	}
}

func main() {
	loadEnv()
	http.HandleFunc("/", rootHandler) // Update this line of code

	log.Println("Time Threshold: ", thresholdTime.String())
	log.Println("Secret URL: ", secretUrl)
	log.Println("Starting server at port http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
