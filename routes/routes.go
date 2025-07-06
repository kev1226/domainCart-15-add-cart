package routes

import (
	"add-service/view"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Habilitar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cart := r.Group("/cart", jwt.AuthGuard("user"))
	{
		cart.POST("", view.AddToCart)
	}
	return r
}
