package repositories

import (
	"altpanel/config"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

// Init function to set collection after DB connection
func NewCustomerRepositoryold() {
	Collection = config.GetCollection("customers")
}

// ---------- Example Methods ----------

// Search customers with skip/limit
func Search(ctx context.Context, filter bson.M, skip int64, limit int64, projection bson.M) ([]bson.M, error) {
	agg := mongo.Pipeline{}

	if len(filter) > 0 {
		agg = append(agg, bson.D{{Key: "$match", Value: filter}})
	} else {
		agg = append(agg, bson.D{{Key: "$sort", Value: bson.M{"_id": -1}}})
	}

	if len(projection) > 0 {
		agg = append(agg, bson.D{{Key: "$project", Value: projection}})
	}

	if skip > 0 {
		agg = append(agg, bson.D{{Key: "$skip", Value: skip}})
	}
	if limit > 0 {
		agg = append(agg, bson.D{{Key: "$limit", Value: limit}})
	}

	cur, err := Collection.Aggregate(ctx, agg, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Count customers
func Count(ctx context.Context, filter bson.M) (int64, error) {
	return Collection.CountDocuments(ctx, filter)
}

func FindOne(ctx context.Context, filter bson.M, projection bson.M) (bson.M, error) {
	opts := options.FindOne().SetProjection(projection)
	var result bson.M
	err := Collection.FindOne(ctx, filter, opts).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return bson.M{}, nil
	}
	return result, err
}

// Update customer
func Update(ctx context.Context, filter bson.M, update bson.M) (int64, error) {
	res, err := Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// Delete customer
func Delete(ctx context.Context, filter bson.M) (int64, error) {
	res, err := Collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// GetSummary example (simplified version)
func GetSummary(ctx context.Context, customerID string) (bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"_id": objID}}},
		{{Key: "$project", Value: bson.M{"phones": 0}}},
	}

	cur, err := Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) > 0 {
		return results[0], nil
	}
	return bson.M{}, nil
}
