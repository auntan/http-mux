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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := api.QueryUrls(r.Context(), urls)
	if response.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		println(err.Error())
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	println(r.URL.String())
	time.Sleep(time.Second / 2)
	//w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(r.URL.RequestURI()); err != nil {
		println(err.Error())
	}
}
