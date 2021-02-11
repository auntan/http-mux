package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"http-mux/internal/api"
	"http-mux/internal/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var urls []string
	err := decoder.Decode(&urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate(urls); err != nil {
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

func validate(urls []string) error {
	if len(urls) > config.MaxUrls {
		return fmt.Errorf("too much urls, %v allowed but %v provided", config.MaxUrls, len(urls))
	}

	if len(urls) == 0 {
		return fmt.Errorf("empty urls list")
	}

	for i := range urls {
		_, err := url.Parse(urls[i])
		if err != nil {
			return fmt.Errorf("invalid url %q: %w", urls[i], err)
		}

	}

	return nil
}
