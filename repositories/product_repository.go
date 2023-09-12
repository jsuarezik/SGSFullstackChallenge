// repositories/product_repository.go

package repositories

import (
	"context"
	"time"

	"sgs_fullstack_challenge/configs"
	"sgs_fullstack_challenge/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	GetAll(page int, size int, sortBy string, sortOrder int, query string) ([]models.Product, int64, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	collection := configs.LoadConfig().MongoCollectionName
	return &productRepository{collection: db.Collection(collection)}
}

func (r *productRepository) GetAll(page int, size int, sortBy string, sortOrder int, query string) ([]models.Product, int64, error) {
	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Pagination
	skip := (page - 1) * size
	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(size))

	//Sorting
	if sortBy != "" {
		options.SetSort(bson.D{{Key: sortBy, Value: sortOrder}})
	}

	filter := bson.M{}
	//Filtering
	if query != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"name": bson.M{"$regex": query, "$options": "i"}},
				{"category": bson.M{"$regex": query, "$options": "i"}},
			},
		}
	}

	cursor, err := r.collection.Find(ctx, filter, options)

	if err != nil {
		return nil, 0, err
	}

	totalCount, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}
