package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/invoices", GetInvoices())
	IncomingRoutes.GET("/invoices/:invoice_id", GetInvoice())
	IncomingRoutes.POST("/invoices", CreateInvoice())
	IncomingRoutes.PATCH("/invoices/:invoice_id", UpdateInvoice())
}
