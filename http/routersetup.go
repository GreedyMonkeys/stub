package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() error {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/endpoint").
		HandlerFunc(getFunction)

	router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(postFunction)

	return http.ListenAndServe(":8080", router)
}
