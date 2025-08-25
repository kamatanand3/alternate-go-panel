package services

import (
	// "fmt"
	"altpanel/models"
	"altpanel/repositories"
)



func GetAllConfig() ([]models.Config, error) {
	repositories.InitConfigRepository()
	configs, err := repositories.GetAllConfig()
	// fmt.Printf("Repo response: %+v\n", configs)
	if err != nil {
		return nil, err
	}
	// Dump the repo response for debugging
	
	return configs, nil
}

