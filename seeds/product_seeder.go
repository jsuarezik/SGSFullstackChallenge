package seeds

import (
	"context"
	"fmt"
	"log"
	"sgs_fullstack_challenge/models"
	"sgs_fullstack_challenge/utils"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProductSeeder(collection *mongo.Collection, number int) {
	for i := 0; i < number; i++ {
		product := models.Product{
			ID:          uuid.New().String(),
			Category:    utils.GetFakeCategory(),
			Description: fake.Paragraph(),
			Discount: models.Discount{
				Status: utils.GetRandomBool(),
				Value:  utils.GetRandomFloat(10, 100),
			},
			IsActive: utils.GetRandomBool(),
			Name:     fake.ProductName(),
			Picture:  fmt.Sprintf("https://picsum.photos/700?random=%d", utils.GetRandomInt(100)),
			Price:    utils.GetRandomFloat(10, 100),
			Stock:    utils.GetRandomInt(50),
		}
		_, err := collection.InsertOne(context.Background(), product)
		if err != nil {
			log.Fatalf("Failed the insert product: %v", err)

		}
	}
	log.Printf("Seeded %v products", number)
}
