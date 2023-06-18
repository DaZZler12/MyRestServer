package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) DeleteItemByName(name string) error {
	filter := bson.M{"item_name": name}
	err := s.store.DeleteItemByName(filter)
	if err != nil {
		return serror.NotFoundError("Error in finding item")
	}
	return nil

}
