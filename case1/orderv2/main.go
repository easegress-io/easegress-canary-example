package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order struct {
	OrderID      string `json:"order_id"`
	RoleID       string `json:"role_id"`
	OrderStatus  int    `json:"order_status"`
	OrderVersion string `json:"order_version"`
}

func main() {
	v2 := gin.Default()
	v2.GET("/order", func(c *gin.Context) {
		order := Order{
			OrderID:      "5245000",
			RoleID:       "44312",
			OrderStatus:  1,
			OrderVersion: "v2",
		}
		c.JSON(http.StatusOK, order)
	})
	v2.Run("0.0.0.0:5002")
}
