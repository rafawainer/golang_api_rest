package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rafawainer/golang_api_rest/api-go-rest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
