package services

import (
	"altpanel/repositories"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDigiScore() (bson.M, error) {
	// Always pass context
	repositories.NewCustomerRepository()
	ctx := context.TODO()

	filter := bson.M{"user_reference_number": "USER9153168152498563"}
	projection := bson.M{
		"user_reference_number": 1,
		"first_name":            1,
		"last_name":             1,
		"email":                 1,
	}

	customer, err := repositories.FindOne(ctx, filter, projection)
	if err != nil {
		return nil, err
	}

	// Debug print
	fmt.Println("✅ Customer found:", customer)

	return customer, nil
}
