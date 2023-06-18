package service

import (
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/store"
	"github.com/DaZZler12/MyRestServer/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	SignIn(email string, password string) (models.User, error)
	SignUp(data models.User) error
	InsertItem(data models.Item) error
	GetAllItems(data utils.Pagination, filters bson.D) ([]models.Item, int64, error)
	GetItemByID(id primitive.ObjectID) (item models.Item, err error)
	GetItemByName(filter bson.M) (models.Item, error)
	UpdateItemByID(data models.Item, id primitive.ObjectID) error
	UpdateItemByName(data models.Item, name string) error
	DeleteItemById(id primitive.ObjectID) error
	DeleteItemByName(name string) error
}

type Service struct {
	store store.Store
}

func NewUserService(store store.Store) *Service {
	return &Service{
		store: store,
	}
}
