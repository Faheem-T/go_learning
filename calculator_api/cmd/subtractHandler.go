package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SubtractBody struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body SubtractBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	result := body.Number1 - body.Number2

	resBody, err := json.Marshal(struct {
		Result int `json:"result"`
	}{Result: result})
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(resBody)

	if writeErr != nil {
		log.Println(writeErr.Error())
	}
}
