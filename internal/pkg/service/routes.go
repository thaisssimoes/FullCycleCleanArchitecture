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

	var orders Order

	db, err := repository.OpenConnection("localhost", "postgres", "password", "postgres", 5432)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	defer repository.CloseConnection(db)

	rows, err := db.Query("Select * from orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	for rows.Next() {
		err = rows.Scan(&orders.ID, &orders.Products, &orders.Total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
		}
	}

	c.IndentedJSON(http.StatusOK, orders)

}
