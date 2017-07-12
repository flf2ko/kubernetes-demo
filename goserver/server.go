package main

import (
	"fmt"
        "io"
        "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive request!")
        io.WriteString(w, "Hello world! This is a go server! version 2")
}

func main() {
        http.HandleFunc("/", hello)
        http.ListenAndServe(":8011", nil)
}
