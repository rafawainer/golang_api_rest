package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafawainer/golang_api_rest/api-go-rest/models"
	"net/http"
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
		c.JSON(http.StatusOK, gin.H{
			"id":     id,
			"status": http.StatusOK,
		})
	})

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
