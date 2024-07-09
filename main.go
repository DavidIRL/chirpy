package main

import (
    "log"
    "net/http"
)

type apiConfig struct {
    fileserverHits int
}


func main() {
    const port = "8080"
    const path = "."

    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir(path)))
    server := &http.Server{
        Addr:   ":" + port,
        Handler:  mux,
    }

    

    log.Printf("Serving files from %s on port: %s\n", path, port)
    log.Fatal(server.ListenAndServe())
}
