package main

import (
	"example/micro/http"

	"log"
)

func main() {
	log.Fatal(http.SetupRouter())
}
