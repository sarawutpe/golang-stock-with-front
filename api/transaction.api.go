package api

import (
	"main/db"
	"main/interceptor"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", interceptor.JwtVerify, getTransaction)
		transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}
}

type TransactionResult struct {
	ID            uint      `json:"id"`
	Total         float64   `json:"total"`
	Paid          float64   `json:"paid"`
	Change        float64   `json:"change"`
	PaymentType   string    `json:"payment_type"`
	PaymentDetail string    `json:"payment_detail"`
	OrderList     string    `json:"order_list"`
	Staff         string    `json:"staff_id"`
	CreatedAt     time.Time `json:"created_at"`
}

func getTransaction(c *gin.Context) {
	var result []TransactionResult
	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_detail, order_list, users.username as Staff, transactions.created_at FROM transactions join users on transactions.staff_id = users.id order by transactions.created_at DESC", nil).Scan(&result)
	c.JSON(200, result)
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(404, gin.H{"result": "nok"})
	}
}
