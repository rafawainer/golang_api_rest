package main

import (
	"fmt"
	"github.com/rafawainer/golang_api_rest/api-go-rest/models"
	"github.com/rafawainer/golang_api_rest/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			Nome:     "Nome 1",
			Historia: "Historia 1",
		},
		{
			Nome:     "Nome 2",
			Historia: "Historia 2",
		},
	}

	fmt.Println("Iniciando o Servidor Rest com Go")
	routes.HandleRequest()
}
