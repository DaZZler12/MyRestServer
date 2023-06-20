package service

import (
	"fmt"

	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) UpdateItemByID(data models.Item, id primitive.ObjectID) error {

	filter := bson.M{"id": id}

	// Build the update statement
	update := bson.M{"$set": bson.M{"brand": data.Brand}}

	err := s.store.UpdateItemByID(filter, update)
	if err != nil {
		return serror.InternalServerError(fmt.Sprintf("error calling store, %v", err))
	}
	return nil

}
