package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers

func checkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	result := fmt.Sprintf("Hello World!!\n")
	w.Write([]byte(result))
}


func testHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	w.Write([]byte("Ok\n"))
}
