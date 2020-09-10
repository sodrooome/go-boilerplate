package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

type Products struct {
	ID uint `gorm:"primary_key:auto_increment" json:"product_id"`
	Name string `gorm:"size:256;not null;unique" json:"name"`
	Price uint `json:"price"`
	Discount uint `json:"discount_product"`
	Stock uint `json:"stock_product"`
	Description string `gorm:"size:500;not null" json:"description"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Cache struct {
	gin.ResponseWriter
	Cache *cache.Cache
	URlPath string
}

type Backoff struct {
	Interval int
}

type Grpc struct {
	Key string
}

type ProductGrpcHandler struct {
	ProductGrpc Products
}