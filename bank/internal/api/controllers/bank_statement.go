package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/controllers/model/response"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (sc *bankControllerInterface) BankStatement(c *gin.Context) {
	logger.Info("Init BankStatement controller",
		zap.String("journey", "Statement"),
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error to get ID from request",
			nil,
			zap.String("journey", "Statement"),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": "ID is required",
		})
		return
	}

	statement, err := sc.service.BankStatement(c.Request.Context(), id)
	if err != nil {
		logger.Error("Error to get bank statement",
			err,
			zap.String("journey", "Statement"),
		)

		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	logger.Info("Bank statement retrieved",
		zap.String("journey", "Statement"),
	)

	c.JSON(http.StatusOK, response.NewBankStatementResponse(statement))
}
