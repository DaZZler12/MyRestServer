package store

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Store interface {
	GetUserByEmail(filter bson.M) (models.User, error)
	InsertUser(data models.User) error
	GetItemByModelBrand(filter bson.M) (models.Item, error)
	GetItemByName(filter bson.M) (models.Item, error)
	GetItemByID(filter bson.M) (models.Item, error)
	InsertItem(data models.Item) error
	GetAllItems(start int, end int, filters bson.D) ([]models.Item, int64, error)
	UpdateItemByID(filter bson.M, updater bson.M) error
	UpdateItemByName(filter bson.M, updater bson.M) error
	DeleteItemById(filter bson.M) error
	DeleteItemByName(filter bson.M) error
	Count(filter bson.D) (int64, error)
}
