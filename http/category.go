package http

import (
	"example/micro/model"
	"example/micro/repository"

	"log"
	"net/http"

	"encoding/json"
)

// post a category
func postCategory(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	category := &model.Category{}
	err := decoder.Decode(category)
	if err != nil {
		panic(err)
	}
	category, err = repository.CreateCategory(category)
	if err != nil {
		panic(err)
	}
	log.Println(category)
}

// get all categories
func getCategory(w http.ResponseWriter, r *http.Request) {

	category_list, err := repository.FindAllCategory()
	if err != nil {
		panic(err)
	}

	payload, err := json.Marshal(category_list)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
