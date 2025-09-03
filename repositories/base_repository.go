// repositories/base_repository.go
package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository[T any] struct {
	Collection *mongo.Collection
}

func (r *BaseRepository[T]) FindOne(ctx context.Context, filter bson.M, projection bson.M) (bson.M, error) {
	opts := options.FindOne().SetProjection(projection)
	var result bson.M
	err := r.Collection.FindOne(ctx, filter, opts).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}
