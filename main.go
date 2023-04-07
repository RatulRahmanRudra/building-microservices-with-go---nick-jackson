package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("error"))
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}

		log.Printf("data %s\n", d)

		fmt.Fprintf(rw, "hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
