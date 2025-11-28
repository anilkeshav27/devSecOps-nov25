package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader((http.StatusOK))
		fmt.Fprint(w, "I am alive and all okay and well also")

		token := jwt.New(jwt.SigningMethodHS256)

		tokenString, err := token.SignedString([]byte("dummy-secret"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("generated token is :", tokenString)

	})

	log.Println("Server is on localhost at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
