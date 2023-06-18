package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) GetItemByID(id primitive.ObjectID) (models.Item, error) {
	filter := bson.M{"id": id}
	item, err := s.store.GetItemByID(filter)
	if err != nil {
		// Handle the error using the spearetaerror package
		return models.Item{}, serror.NotFoundError("Error in finding item")
	}
	return item, nil

}
