package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AddBody struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body AddBody

	// body, err := io.ReadAll(r.Body)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON ", http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	result := body.Number1 + body.Number2

	fmt.Fprintf(w, "Result %v", result)
}
