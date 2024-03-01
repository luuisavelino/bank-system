package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/controllers/model/request"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (sc *bankControllerInterface) BankTransaction(c *gin.Context) {
	logger.Info("Init BankTransaction controller",
		zap.String("journey", "Transaction"),
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error to get UUID from request",
			err,
			zap.String("journey", "Transaction"),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	var bankTransactionRequest request.BankTransactionRequest
	if err := c.ShouldBindJSON(&bankTransactionRequest); err != nil {
		logger.Error("Error to bind json",
			err,
			zap.String("journey", "Transaction"),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	bankTransaction := models.NewBankTransactionDomain(
		bankTransactionRequest.Valor,
		bankTransactionRequest.Descricao,
		bankTransactionRequest.Tipo,
	)

	err = sc.service.BankTransaction(c.Request.Context(), id, bankTransaction)
	if err != nil {
		logger.Error("Error to create bank transaction",
			err,
			zap.String("journey", "Transaction"),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	logger.Info("Transaction done with success",
		zap.String("journey", "Transaction"),
	)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success", "message": "transaction done with success",
	})
}
