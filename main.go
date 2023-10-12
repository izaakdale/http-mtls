package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	if os.Getenv("ENABLE_TLS") == "true" {
		log.Printf("Serving https on %+v\n", addr)
		http.ListenAndServeTLS(addr, os.Getenv("SERVER_CRT"), os.Getenv("SERVER_KEY"), mux)
	} else {
		log.Printf("Serving http on %+v\n", addr)
		http.ListenAndServe(addr, mux)
	}
}
