package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte(`{"code":0,"data":"Hello Gopher ! ! !"}`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
