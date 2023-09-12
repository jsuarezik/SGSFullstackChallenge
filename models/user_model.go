package models

type Product struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	Category    string   `bson:"category" json:"category"`
	Description string   `bson:"description" json:"description"`
	Discount    Discount `bson:"discount" json:"discount"`
	IsActive    bool     `bson:"isActive" json:"isActive"`
	Name        string   `bson:"name" json:"name"`
	Picture     string   `bson:"picture" json:"picture"`
	Price       float64  `bson:"price" json:"price"`
	Stock       int      `bson:"stock" json:"stock"`
}

type Discount struct {
	Status bool    `bson:"status" json:"status"`
	Value  float64 `bson:"value" json:"value"`
}
