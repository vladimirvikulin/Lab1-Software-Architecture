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
