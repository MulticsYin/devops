package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"code":0,"data":"successful"}`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
