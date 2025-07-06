package presenter

import (
	"add-service/dto"
	"add-service/entity"
	"add-service/external"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func AddToCart(dto dto.AddToCartDTO, userID interface{}, redisClient *redis.Client) error {
	// Validar que el productId sea un número válido
	if _, err := strconv.Atoi(dto.ProductID); err != nil {
		return fmt.Errorf("ID de producto inválido")
	}

	// Obtener el producto desde el microservicio externo
	product, err := external.GetProductByID(dto.ProductID)
	if err != nil {
		return err
	}

	// Validar stock
	if dto.Quantity > product.Stock {
		return fmt.Errorf("stock insuficiente (%d disponibles)", product.Stock)
	}

	// Crear y guardar en Redis
	item := entity.CartItem{
		ProductID: fmt.Sprintf("%d", product.ID),
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  dto.Quantity,
	}
	itemJSON, _ := json.Marshal(item)
	key := fmt.Sprintf("cart:%v", userID)

	return redisClient.HSet(ctx, key, item.ProductID, itemJSON).Err()
}
