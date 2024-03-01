package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/controllers"
)

func InitRoutes(r *gin.RouterGroup, controller controllers.BankControllerInterface) {
	v1 := r.Group("/clientes")
	{
		v1.GET("/:id/extrato", controller.BankStatement)
		v1.POST("/:id/transacoes", controller.BankTransaction)
	}
}
