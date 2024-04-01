package routes

import "github.com/gin-gonic/gin"

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", GetOrderItem())
	incomingRoutes.GET(".orderItems-order/:order_id", GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", UpdateOrderItem())
}
