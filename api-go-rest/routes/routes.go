package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rafawainer/golang_api_rest/api-go-rest/models"
	"net/http"
	"strconv"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Home Page")
	})

	r.GET("/api/personalidades", func(c *gin.Context) {
		p := models.Personalidades
		c.JSON(http.StatusOK, gin.H{
			"response": p,
			"status":   http.StatusOK,
		})
	})

	r.GET("/api/personalidades/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		naoNumerico := idInt < 1
		if err != nil || naoNumerico {
			message := "Erro para converter o id " + id + " para numerico"
			fmt.Println(message)
			c.JSON(http.StatusBadRequest, gin.H{
				"id":      id,
				"status":  http.StatusBadRequest,
				"message": message,
			})
		} else {
			p := models.Personalidades
			r, message := models.RetornaUmaPersonalidade(p, idInt)
			if r.Id > 0 {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"data":    r,
					"status":  http.StatusOK,
					"message": message,
				},
				)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"status":  http.StatusOK,
					"message": message,
				},
				)
			}
		}
	},
	)

	//r.http.HandleFunc(
	//	"/",
	//	controllers.Home,
	//)
	//http.HandleFunc(
	//	"/api/personalidades",
	//	controllers.TodasPersonalidades,
	//)
	r.Run(":8000")
	//log.Fatal(http.ListenAndServe(":8000", nil))
}
