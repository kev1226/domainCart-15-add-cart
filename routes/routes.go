package routes

import (
	"add-service/view"

	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	cart := r.Group("/cart", jwt.AuthGuard("user"))
	{
		cart.POST("", view.AddToCart)
	}
	return r
}
