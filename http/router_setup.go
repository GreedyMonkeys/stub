package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() error {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/category").
		HandlerFunc(getCategory)

	router.
		Methods("POST").
		Path("/category").
		HandlerFunc(postCategory)

	return http.ListenAndServe(":8080", router)
}
