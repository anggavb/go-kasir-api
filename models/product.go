package models

import "database/sql"

type Product struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Price      int           `json:"price"`
	Stock      int           `json:"stock"`
	CategoryID sql.NullInt64 `json:"category_id"`
	Category   Category
}
