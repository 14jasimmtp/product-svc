package models

type StockDecreaseLog struct {
	ID           int64 `json:"id" gorm:"primaryKey"`
	OrderId      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_id"`
}

type Product struct {
	ID                int64            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Price             int64            `json:"price"`
	Stock             int64            `json:"stock"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
