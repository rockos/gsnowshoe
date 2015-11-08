package main

import (
	. "github.com/rockos/gap"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3013", router))
}
