package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// DefaultPort default for http server listen port
const DefaultPort = ":8080"

// OmikujiResponse omikuji-api response
type OmikujiResponse struct {
	ResultCode int    `json:"result_code"`
	Result     string `json:"result"`
}

func omikuji() (int, string) {
	t := time.Now()
	if t.Month() == time.January && t.Day() >= 1 && t.Day() <= 3 {
		return 0, "daikichi"
	}
	rand.Seed(t.UnixNano())
	s := rand.Intn(6) + 1

	switch s {
	case 1:
		return s, "kyo"
	case 2, 3:
		return s, "kichi"
	case 4, 5:
		return s, "chukichi"
	case 6:
		return s, "daikichi"
	default:
		return s, ""
	}
}

// HandleOmikujiAPI omikuji api
func HandleOmikujiAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var v OmikujiResponse
	v.ResultCode, v.Result = omikuji()

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("error:", err)
	}
}

func main() {
	port := DefaultPort
	http.HandleFunc("/", HandleOmikujiAPI)
	log.Fatal(http.ListenAndServe(port, nil))
}
