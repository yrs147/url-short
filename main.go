package main

import (
	"fmt"
	"net/http"

	"github.com/yrs147/url-short/routes"
)

func main() {
	routes.SetupRoutes()

	port := 9090
	fmt.Printf("Server running on %d ...\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d",port),nil)
}
