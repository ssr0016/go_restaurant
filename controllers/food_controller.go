package controller

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Foods")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get Food")
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Create Food")
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Update Food")
	}
}
