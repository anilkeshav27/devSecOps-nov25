package main

import (
	"fmt"
	"log"
	"net/http"
)

// 

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader((http.StatusOK))
		fmt.Fprint(w, "I am alive and all okay and well also")
	})

	log.Println("Server is on localhost at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
