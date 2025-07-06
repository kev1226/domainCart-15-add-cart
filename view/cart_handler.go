package view

import (
	"add-service/config"
	"add-service/dto"
	"add-service/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

var RedisClient = config.InitRedis()

func AddToCart(c *gin.Context) {
	var input dto.AddToCartDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	err := presenter.AddToCart(input, userID, RedisClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto agregado al carrito"})
}
