package external

import (
	"add-service/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

const productServiceURL = "http://44.216.156.192:3022/products/"

func GetProductByID(productID string) (*entity.ProductResponse, error) {
	resp, err := http.Get(productServiceURL + productID)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con el servicio de productos")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("producto no encontrado (status %d)", resp.StatusCode)
	}

	var product entity.ProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("error al procesar la respuesta del producto")
	}

	return &product, nil
}
