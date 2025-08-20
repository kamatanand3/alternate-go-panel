package services

import (
	"altpanel/models"
	"altpanel/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (primitive.ObjectID, error) {
	return repositories.CreateUser(user)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserByID(id primitive.ObjectID) (models.User, error) {
	return repositories.GetUserByID(id)
}

func UpdateUser(id primitive.ObjectID, user models.User) error {
	return repositories.UpdateUser(id, user)
}

func DeleteUser(id primitive.ObjectID) error {
	return repositories.DeleteUser(id)
}
