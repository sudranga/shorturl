package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func initHttpServer() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8000", nil)
}
