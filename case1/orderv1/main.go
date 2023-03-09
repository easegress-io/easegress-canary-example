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
	v1.GET("/order", func(c *gin.Context) {
		order := Order{
			OrderID:      "5245000",
			RoleID:       "44312",
			OrderStatus:  1,
			OrderVersion: "v1",
		}
		c.JSON(http.StatusOK, order)
	})
	v1.Run("0.0.0.0:5001")
}
