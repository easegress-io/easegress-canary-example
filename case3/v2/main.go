package main

import (
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"net/http"
)

type Order struct {
	OrderID      string `json:"order_id"`
	RoleID       string `json:"role_id"`
	OrderStatus  int    `json:"order_status"`
	OrderVersion string `json:"order_version"`
}

type Notify struct {
	Order
	NotifyStatus  int    `json:"notify_status"`
	NotifyVersion string `json:"notify_version"`
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
		resp := sendNotify(order, c.Request)
		notify := Notify{}
		resp.JSON(&notify)
		c.JSON(resp.StatusCode, notify)
	})
	v2.POST("/notify", func(c *gin.Context) {
		order := Order{}
		c.Bind(&order)
		c.JSON(http.StatusOK, gin.H{
			"order_id": order.OrderID, "role_id": order.RoleID, "order_status": order.OrderStatus, "notify_status": 1, "notify_version": "v2", "order_version": order.OrderVersion,
		})
	})
	v2.Run("0.0.0.0:5002")
}

func sendNotify(order Order, request *http.Request) *grequests.Response {
	options := &grequests.RequestOptions{
		JSON: order,
	}
	devices, ok := request.Header["X-Ua-Device"]
	if ok {
		options.Headers = map[string]string{
			"X-Ua-Device": devices[0],
		}
	}
	system, ok := request.Header["X-Ua-Os"]
	if ok {
		options.Headers = map[string]string{
			"X-Ua-Os": system[0],
		}
	}
	resp, _ := grequests.Post("http://0.0.0.0:8888/notify", options)
	return resp
}
