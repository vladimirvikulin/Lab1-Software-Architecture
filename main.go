package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type TimeResponse struct {
    Time string json:"time"
}

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" && r.URL.Path == "/time" {
            currentTime := time.Now().Format(time.RFC3339)
            response := TimeResponse{Time: currentTime}
            json.NewEncoder(w).Encode(response)
        } else {
            http.NotFound(w, r)
        }
    }
}
