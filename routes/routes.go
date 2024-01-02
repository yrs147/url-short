package routes

import (
	"net/http"
	"github.com/yrs147/url-short/controllers"
	
)

func SetupRoutes(){
	http.HandleFunc("api/v1/shorten",controllers.ShortenUrlHandler)
	http.HandleFunc("api/v1/shorurl",controllers.GetShortHandler)
}