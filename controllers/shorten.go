package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/yrs147/url-short/model"
)

var urlMappings map[string]string

func init() {
	urlMappings = make(map[string]string)
}

func ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {

	var req model.Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	shortUrl := generateURL()

	urlMappings[shortUrl] = req.URL

	response := model.Response{
		URL: shortUrl,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func generateURL() string {
	uuid := uuid.New()

	id := uuid.String()[:8]

	return id

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	longURL, exists := urlMappings[shortURL]
	if exists {
		http.Redirect(w, r, longURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}
