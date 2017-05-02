package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func response(rw http.ResponseWriter, request *http.Request) {
	json, err := json.Marshal("svc one alive!")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rw.Write([]byte(json))
}

func main() {
	log.Println("SVC One starting on :8000")
	http.HandleFunc("/", response)
	http.ListenAndServe(":8000", nil)
}
