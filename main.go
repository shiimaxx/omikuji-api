package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// HandleOmikujiAPI omikuji api
func HandleOmikujiAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Result string `json:"result"`
	}{}
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(5) + 1

	switch s {
	case 1:
		v.Result = "kyo"
	case 2, 3:
		v.Result = "kichi"
	case 4, 5:
		v.Result = "chukichi"
	case 6:
		v.Result = "daikichi"
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("error:", err)
	}
}

func main() {
	http.HandleFunc("/", HandleOmikujiAPI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
