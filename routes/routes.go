package routes

import (
	"net/http"
	"github.com/yrs147/url-short/controllers"
	
)

func SetupRoutes(){
	http.HandleFunc("/shorten",controllers.ShortenUrlHandler)
	http.HandleFunc("/",controllers.Redirect)
}