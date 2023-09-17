package main

import (
	"fmt"
	models "github.com/rafawainer/go_api_est/models"
	routes "github.com/rafawainer/go_api_est/routes"
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
