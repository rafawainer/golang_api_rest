package routes

import (
	controllers "github.com/rafawainer/go_api_est/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc(
		"/",
		controllers.Home,
	)
	http.HandleFunc(
		"/api/personalidades",
		controllers.TodasPersonalidades,
	)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
