package http

import (
	"example/micro/model"

	"log"
	"net/http"

	"encoding/json"
)

func postFunction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var suka model.Suka
	err := decoder.Decode(&suka)
	if err != nil {
		panic(err)
	}
	log.Println(suka)
}

func getFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing")
	suka := &model.Suka{Name: "Suka", Surname: "Suka2"}

	payload, err := json.Marshal(suka)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
