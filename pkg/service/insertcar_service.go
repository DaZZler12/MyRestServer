package service

import (
	"time"

	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) InsertItem(data models.Item) error {

	filter := bson.M{"item_name": data.Item_Name}
	_, err := s.store.GetItemByModelBrand(filter)
	if err == nil {
		return serror.AlreadyInUse("Item Details Already there")
	}
	data.CreatedAt = time.Now()
	// data.UpdatedAt = time.Now()
	err = s.store.InsertItem(data)
	if err != nil {
		return serror.BadRequestError("Error adding the Item")
	}

	return nil

}
