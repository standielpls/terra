package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	type helloRes struct {
		Message string `json:"message"`
		Error   struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		b, _ := json.Marshal(helloRes{Message: "hello hello"})
		w.Write(b)
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("can't serve on 8000: %s", err.Error())
	}
	log.Println("serving on :8080")
}
