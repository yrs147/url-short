package controllers

import (
	"encoding/json"
	"net/http"

	b64 "encoding/base64"

	"github.com/google/uuid"
	"github.com/yrs147/url-short/model"
)


func ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	
	var req model.Request
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err !=nil{
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	shortUrl := generateURL()

	response := model.Response{
		URL : shortUrl,
	}

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(response)

}


func generateURL() string{
	uuid := uuid.New()

	id :=  uuid.String()[:8]

	hashid := b64.StdEncoding.EncodeToString([]byte(id))

	return hashid[:5]
}