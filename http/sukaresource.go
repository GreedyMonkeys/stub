package http

import (
	"example/micro/model"
	"example/micro/repository"

	"log"
	"net/http"

	"encoding/json"
)

// post a suka
func postFunction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var suka model.Suka
	err := decoder.Decode(&suka)
	if err != nil {
		panic(err)
	}
	suka, err = repository.CreateSuka(suka)
	log.Println(suka)
}

// get all sukas
func getFunction(w http.ResponseWriter, r *http.Request) {
	// suka := &model.Suka{Name: "Suka", Surname: "Suka2"}

	sukas, err := repository.FindAllSuka()

	payload, err := json.Marshal(sukas)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
