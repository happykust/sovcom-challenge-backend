package http

import (
	"api-gateway/internal/app/api/domain/payments/transaction"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	transactionGroup := route.Group("/transaction")
	transactionGroup.POST("/create", transaction.CreateTransaction)
	transactionGroup.GET("/get", transaction.GetTransaction)
	transactionGroup.POST("/update", transaction.UpdateTransactionStatus)
}
