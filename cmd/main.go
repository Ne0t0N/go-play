package main

import (
	"flag"
	"fmt"
	"go-micro-book/config"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 80, "port on which http server should listen")
	flag.Parse()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config, error: %v\n", err.Error())
	}
	log.Printf("app config: %+v\n", cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	})

	log.Printf("app is running on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
