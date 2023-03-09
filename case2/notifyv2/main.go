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
	v1 := gin.Default()
	v1.POST("/notify", func(c *gin.Context) {
		order := Order{}
		c.Bind(&order)
		c.JSON(http.StatusOK, gin.H{
			"order_id": order.OrderID, "role_id": order.RoleID, "order_status": order.OrderStatus, "notify_status": 1, "notify_version": "v2", "order_version": order.OrderVersion,
		})
	})
	v1.Run("0.0.0.0:5002")
}
