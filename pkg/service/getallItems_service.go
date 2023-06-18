package service

import (
	"fmt"

	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) GetAllItems(start int, end int, filter bson.D) ([]models.Item, int64, error) {
	items, total, err := s.store.GetAllItems(start, end, filter)
	if err != nil {
		fmt.Print(err)
		return nil, 0, serror.InternalServerError("Failed to Retrieve the items")
	}

	return items, total, nil

}

func (s *Service) Count(filter bson.D) (int64, error) {
	count, err := s.store.Count(filter)
	if err != nil {
		return 0, err
	}
	return count, err
}
