package service

import (
	"FullCycleCleanArchitecture/internal/pkg/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	var order Order
	var orders []Order

	db, err := repository.OpenConnection("localhost", "postgres", "password", "postgres", 5432)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	defer repository.CloseConnection(db)

	rows, err := db.Queryx("Select * from orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	for rows.Next() {
		err = rows.StructScan(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
		}
		orders = append(orders, order)
	}

	c.IndentedJSON(http.StatusOK, orders)

}
