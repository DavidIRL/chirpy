package main

import "net.http"


func (cfg *apiConfig) handlerReset(write http.ResponseWriter, req *http.Request) {
    cfg.fileserverHits = 0
    write.WriteHeader(http.StatusOK)
    write.Write([]byte("Hits have been reset to 0"))
}
