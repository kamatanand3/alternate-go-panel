package models

type Customer struct {
	ID                  string `bson:"_id,omitempty" json:"id"`
	FirstName           string `bson:"first_name" json:"first_name"`
	LastName            string `bson:"last_name" json:"last_name"`
	Email               string `bson:"email" json:"email"`
	UserReferenceNumber string `bson:"user_reference_number" json:"user_reference_number"`
}
