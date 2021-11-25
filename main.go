package main

import (
	"example/micro/http"
	"example/micro/repository"

	"log"
)

func main() {
	repository.CreateDbConnection()
	log.Fatal(http.SetupRouter())
}
