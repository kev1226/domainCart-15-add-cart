package kafka

type ProductRequest struct {
	RequestID string `json:"requestId"`
	ProductID string `json:"productId"`
}

type ProductResponse struct {
	RequestID   string      `json:"requestId"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Description string      `json:"description"`
	Stock       int         `json:"stock"`
	SKU         string      `json:"sku"`
	IsPublished bool        `json:"isPublished"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
	DeletedAt   interface{} `json:"deletedAt"`
}
