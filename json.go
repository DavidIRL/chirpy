package main

import (
    "encoding/json"
    "log"
    "net/http"
)


func respondWithError(write http.ResponseWriter, code int, msg string) {
    if code > 499 {
        log.Printf("Responding with 5XX error: %s", msg)
    }
    type errorResponse struct {
        Error string `json:"error"`
    }
    respondWithJSON(write, code, errorResponse {
        Error: msg,
    })
}


func respondWithJSON(write http.ResponseWriter, code int, payload inteface{}) {
    write.Header().Set("Content-Type", "application/json")
    data, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling JSON: %s", err)
        write.WriteHeader(500)
        return
    }
    write.WriteHeader(code)
    write.Write(data)
}
