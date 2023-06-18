package service

import (
	"fmt"
	"time"

	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) UpdateItemByName(data models.Item, name string) error {

	filter := bson.M{"item_name": name}

	// Build the update statement
	update := bson.M{"$set": bson.M{"brand": data.Brand, "model": data.Model, "year": data.Year, "price": data.Price, "updated_at": time.Now()}}

	err := s.store.UpdateItemByName(filter, update)
	if err != nil {
		return serror.InternalServerError(fmt.Sprintf("error calling store, %v", err))
	}
	return nil
}
