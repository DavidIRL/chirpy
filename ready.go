package main

import "net/http"


func handlerReadiness(write http.ResponseWriter, req *http.Request) {
    write.Header().Add("Content-Type", "text/plain; charset=utf-8")
    write.WriteHeader(http.StatusOK)
    write.Write([]byte(http.StatusText(http.StatusOK)))
}
