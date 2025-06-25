package presenter

import (
	"add-service/dto"
	"add-service/entity"
	"add-service/kafka"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func AddToCart(dto dto.AddToCartDTO, userID interface{}, redisClient *redis.Client) error {
	// Validar que el productId sea un número válido
	if _, err := strconv.Atoi(dto.ProductID); err != nil {
		return fmt.Errorf("ID de producto inválido")
	}

	// ✅ Generar un requestId único para esta solicitud
	requestID := uuid.New().String()

	req := kafka.ProductRequest{
		RequestID: requestID,
		ProductID: dto.ProductID,
	}
	data, _ := json.Marshal(req)
	if err := kafka.SendMessage(kafka.ProductRequestTopic, requestID, data); err != nil {
		return fmt.Errorf("error solicitando producto")
	}

	// Esperar respuesta filtrando por requestId
	prodRes, err := kafka.WaitForProductResponse(requestID)
	if err != nil {
		return fmt.Errorf("producto no encontrado o tiempo agotado")
	}

	// Verificar stock disponible
	if dto.Quantity > prodRes.Stock {
		return fmt.Errorf("stock insuficiente (%d disponibles)", prodRes.Stock)
	}

	// Crear item y guardarlo en Redis
	item := entity.CartItem{
		ProductID: fmt.Sprintf("%d", prodRes.ID),
		Name:      prodRes.Name,
		Price:     prodRes.Price,
		Quantity:  dto.Quantity,
	}
	itemJSON, _ := json.Marshal(item)
	key := fmt.Sprintf("cart:%v", userID)

	return redisClient.HSet(ctx, key, item.ProductID, itemJSON).Err()
}
