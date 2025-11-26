package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleSum(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body []int
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var result int = 0
	for _, v := range body {
		result += v
	}

	resBody, jsonErr := json.Marshal(struct {
		Result int `json:"result"`
	}{Result: result})
	if jsonErr != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Println(jsonErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(resBody)

	if writeErr != nil {
		log.Println(writeErr.Error())
	}
}
