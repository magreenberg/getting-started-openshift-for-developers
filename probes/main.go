package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	listenAddrDefaultValue = "0.0.0.0:8080"
)

var (
	livenessResponse  = http.StatusServiceUnavailable
	readinessResponse = http.StatusServiceUnavailable
	startupResponse   = http.StatusServiceUnavailable
	startupCount      = 0
	readinessCount    = 0
	livenessCount     = 0
	ready             = false
)

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	log.Println("Hello, world")
}

func startup(w http.ResponseWriter, r *http.Request) {
	startupCount++
	if startupCount == 5 {
		startupResponse = http.StatusOK
		log.Println("Setting startup response to OK.")
	}
	w.WriteHeader(startupResponse)
	log.Println("startup returned: ", startupResponse)
}

func liveness(w http.ResponseWriter, r *http.Request) {
	livenessCount++
	w.WriteHeader(livenessResponse)
	log.Println("liveness returned: ", livenessResponse)
}

func readiness(w http.ResponseWriter, r *http.Request) {
	readinessCount++
	if readinessCount == 4 {
		readinessResponse = http.StatusOK
		log.Println("Setting readiness response to OK. Service should now be available for use.")
		// if readinessResponse == http.StatusOK {
		// 	readinessResponse = http.StatusServiceUnavailable
		// 	log.Println("Setting readiness response to Unavailable.")
		// } else {
		// 	readinessResponse = http.StatusOK
		// 	log.Println("Setting readiness response to OK. Service should now be available for use.")
		// }
		// readinessCount = 0
	}
	w.WriteHeader(readinessResponse)
	log.Println("readiness returned: ", readinessResponse)
}

func toggleReadiness(w http.ResponseWriter, r *http.Request) {
	if ready {
		readinessResponse = http.StatusServiceUnavailable
		livenessResponse = http.StatusServiceUnavailable
		ready = false
	} else {
		readinessResponse = http.StatusOK
		livenessResponse = http.StatusOK
		ready = true
	}
	w.WriteHeader(http.StatusOK)
	log.Println("toggleReadiness returned: ", http.StatusOK)
}


func notifyExit() {
	log.Println("Interrupt caught. Application exiting after liveness probe failures: ", livenessCount)
}

func main() {
	// catch restart
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		notifyExit()
		os.Exit(1)
	}()

	http.HandleFunc("/", serve)
	http.HandleFunc("/startup", startup)
	http.HandleFunc("/liveness", liveness)
	http.HandleFunc("/readiness", readiness)
	http.HandleFunc("/togglereadiness", toggleReadiness)

	log.Println("listening at:", listenAddrDefaultValue)
	http.ListenAndServe(listenAddrDefaultValue, nil)
}
