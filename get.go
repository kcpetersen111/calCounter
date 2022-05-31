package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleGetAll(w http.ResponseWriter, r *http.Request){
	log.Println("Sending All data to ",r.RemoteAddr)
	json.NewEncoder(w).Encode(data)
}

func handleGetLatestDay(w http.ResponseWriter, r *http.Request){
	log.Println("Sending latest day data to ",r.RemoteAddr)
	json.NewEncoder(w).Encode(data[len(data)-1])
}