package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! Nice to meet you")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/**
git pull
go run main.go
git add .
git commit -m "."
git push
*/
