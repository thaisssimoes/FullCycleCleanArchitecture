package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func App() {
	s := gin.Default()
	rotas(s)
	log.Fatalln(s.Run(":8080"))
}

func rotas(s *gin.Engine) {
	s.GET("/order", getAllOrders)
	s.POST("/order", setOrder)
}

func setOrder(c *gin.Context) {}

func getAllOrders(c *gin.Context) {

	var err error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	c.IndentedJSON(http.StatusOK, "mensagem ok")

}
