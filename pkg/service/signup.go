package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) SignUp(data models.User) error {

	existingUser := bson.M{"email": data.Email}
	_, err := s.store.GetUserByEmail(existingUser)
	if err == nil {
		return serror.AlreadyInUse("Email is already in use")
	}
	err = s.store.InsertUser(data)
	if err != nil {
		return serror.BadRequestError("Failed to create an account")
	}

	return nil

}
