package main

import (
	"log"
	"net/http"
	"os"
	"working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.Flags())
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/", hh)

	http.ListenAndServe(":9090", sm)
}
