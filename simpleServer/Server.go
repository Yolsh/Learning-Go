package simpleServer

import (
	"fmt"
	"log"
	"net/http"
)

func ServerStart() {
	http.Handle("/", http.FileServer(http.Dir("./simpleServer/static")))

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}