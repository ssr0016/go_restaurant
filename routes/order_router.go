package routes

import "github.com/gin-gonic/gin"

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", GetOrders())
	incomingRoutes.GET("/orders/:order_id", GetOrder())
	incomingRoutes.POST("/orders", CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", UpdateOrder())
}
