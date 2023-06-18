package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) GetItemByName(filter primitive.M) (models.Item, error) {

	item, err := s.store.GetItemByName(filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Item{}, serror.NotFoundError("Item not found")
		}
		// Handle the error using the spearetaerror package
		return models.Item{}, serror.NotFoundError("Error in finding item")
	}
	return item, nil

}
