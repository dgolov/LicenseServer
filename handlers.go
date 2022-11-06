package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers

func checkHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	result := fmt.Sprintf("Hello World!!\n")
	writer.Write([]byte(result))
}

func testHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	writer.Write([]byte("Ok\n"))
}
