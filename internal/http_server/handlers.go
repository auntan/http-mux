package http_server

import (
	"encoding/json"
	"http-mux/internal/api"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var urls []string
	err := decoder.Decode(&urls)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	response := api.QueryUrls(r.Context(), urls)
	if response.Error != nil {
		w.WriteHeader(500)
	}

	_ = json.NewEncoder(w).Encode(response)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	println(r.URL.String())
	time.Sleep(time.Second / 2)
	//w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(r.URL.RequestURI())
}
