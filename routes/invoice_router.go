package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/invoices", controller.GetInvoices())
	IncomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	IncomingRoutes.POST("/invoices", controller.CreateInvoice())
	IncomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
