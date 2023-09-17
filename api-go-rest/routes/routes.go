package routes

import (
	"github.com/rafawainer/golang_api_rest/api-go-rest/controllers"
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
