package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MultiplyBody struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body MultiplyBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	result := body.Number1 * body.Number2
	resBody, jsonErr := json.Marshal(struct {
		Result int `json:"result"`
	}{Result: result})

	if jsonErr != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Println(jsonErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(resBody)
	if writeErr != nil {
		log.Println(writeErr.Error())
	}
}
