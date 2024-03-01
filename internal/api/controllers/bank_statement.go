package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (sc *bankControllerInterface) BankStatement(c *gin.Context) {
	logger.Info("Init EditBank controller",
		zap.String("journey", "UpdateBank"),
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error to get ID from request",
			nil,
			zap.String("journey", "UpdateBank"),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": "ID is required",
		})
		return
	}

	err = sc.service.BankStatement(c.Request.Context(), id)
	if err != nil {
		logger.Error("Error to update bank",
			err,
			zap.String("journey", "UpdateBank"),
		)

		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	logger.Info("Bank edited with success",
		zap.String("journey", "EditBank"),
	)

	c.JSON(http.StatusOK, gin.H{
		"status": "success", "message": "bank deleted",
	})
}
