package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignIn(usremail string, usrpassword string) (models.User, error) {
	existingUser := bson.M{"email": usremail}
	user, err := s.store.GetUserByEmail(existingUser)
	if err != nil {
		// Handle the error using the spearetaerror package
		return models.User{}, serror.NotFoundError("Error in finding user")
	}
	if user.Email == "" {
		return models.User{}, serror.NotFoundError("User not found")
	}
	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usrpassword))
	if err != nil {
		return models.User{}, serror.UnauthorizedError("fail to login")
	}
	return user, nil

}
