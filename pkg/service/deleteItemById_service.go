package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) DeleteItemById(id primitive.ObjectID) error {
	filter := bson.M{"id": id}
	err := s.store.DeleteItemById(filter)
	if err != nil {
		// Handle the error using the spearetaerror package
		return serror.NotFoundError("Error in finding item")
	}
	return nil

}
