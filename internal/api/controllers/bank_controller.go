package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/service"
)

type BankControllerInterface interface {
	BankStatement(c *gin.Context)
	BankTransaction(c *gin.Context)
}

type bankControllerInterface struct {
	service service.BankServiceInterface
}

func NewBankControllerInterface(serviceInterface service.BankServiceInterface) BankControllerInterface {
	return &bankControllerInterface{
		service: serviceInterface,
	}
}
