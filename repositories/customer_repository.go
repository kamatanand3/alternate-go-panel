package repositories

import (
	"altpanel/config"
	"altpanel/models"
)

type CustomerRepository struct {
	*BaseRepository[models.Customer]
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		BaseRepository: &BaseRepository[models.Customer]{
			Collection: config.GetCollection("customers"),
		},
	}
}
