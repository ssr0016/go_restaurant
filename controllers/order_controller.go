package controller

import "github.com/gin-gonic/gin"

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Invoices")
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Invoices")
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Invoices")
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Invoices")
	}
}
