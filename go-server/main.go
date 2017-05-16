package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type jsonResponse map[string]interface{}

func (r jsonResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		s := ""
		return s
	}
	s := string(b)
	return s
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonResponse{"success": true, "message": "Hello!"})
}

func main() {
	port := os.Getenv("FAKE_SERVER_PORT")
	http.HandleFunc("/", rootHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
