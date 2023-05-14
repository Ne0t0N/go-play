package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 80, "port on which http server should listen")
	addr := fmt.Sprintf(":%d", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	})

	log.Printf("app is running on port %d\n", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
