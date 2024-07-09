package main

import (
    "fmt"
    "net/http"
)


func (cfg *apiConfig) handlerMetrics(write http.ResponseWriter, req *http.Request) {
    write.Header().Add("Content-Type", "text/html")
    write.WriteHeader(http.StatusOK)
    write.Write([]byte(fmt.Sprintf(`
<html>

<body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
</body>

</html>

    `, cfg.fileserverHits)))
}


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
    return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
        cfg.fileserverHits++
        next.ServeHTTP(write, req)
    })
}
