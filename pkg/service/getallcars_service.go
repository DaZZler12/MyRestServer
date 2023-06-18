package service

import (
	"fmt"

	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"github.com/DaZZler12/MyRestServer/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) GetAllItems(pagination utils.Pagination, filter bson.D) ([]models.Item, int64, error) {

	items, total, err := s.store.GetAllItems(pagination, filter)
	if err != nil {
		fmt.Print(err)
		return nil, 0, serror.InternalServerError("Failed to Retrieve the items")
	}

	return items, total, nil

}
