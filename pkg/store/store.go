package store

import (
	"context"
	"log"
	"sync"

	"github.com/DaZZler12/MyRestServer/pkg/config"
	"github.com/DaZZler12/MyRestServer/pkg/database"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type dbStore struct {
	db *mongo.Database
}

var instance Store = (*dbStore)(nil)
var once sync.Once

// GetStore returns the singleton instance of the Store.
func GetStore(cfg config.DatabaseConfig) Store {
	once.Do(func() {
		db, err := database.ConnectToMongoDB(cfg)
		if err != nil {
			log.Fatal(err)
		}
		instance = &dbStore{
			db: db,
		}
	})
	return instance
}

// GetUserByEmail retrieves a user by email.
func (s *dbStore) GetUserByEmail(filter bson.M) (user models.User, err error) {
	collection := s.db.Collection("users")
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

// InsertUser inserts a new user into the database.
func (s *dbStore) InsertUser(data models.User) error {
	collection := s.db.Collection("users")
	id := primitive.NewObjectID()
	user := models.User{
		ID:       id,
		Name:     data.Name,
		Country:  data.Country,
		Email:    data.Email,
		Password: data.Password,
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	// Store the user in the database
	user.Password = string(hashedPassword)
	_, err = collection.InsertOne(context.Background(), user)

	return err
}

// GetItemByModelBrand retrieves a item by model and brand.
func (s *dbStore) GetItemByModelBrand(filter bson.M) (item models.Item, err error) {
	collection := s.db.Collection("item")
	err = collection.FindOne(context.Background(), filter).Decode(&item)
	return item, err
}

// GetItemByName retrieves a item by item_name.
func (s *dbStore) GetItemByName(filter bson.M) (item models.Item, err error) {
	collection := s.db.Collection("item")
	err = collection.FindOne(context.Background(), filter).Decode(&item)
	return item, err
}

// GetItemByID retrieves a item by ID.
func (s *dbStore) GetItemByID(filter bson.M) (item models.Item, err error) {
	collection := s.db.Collection("item")
	err = collection.FindOne(context.Background(), filter).Decode(&item)
	return item, err
}

// InsertItem inserts a new item into database
func (s *dbStore) InsertItem(data models.Item) error {
	collection := s.db.Collection("item")
	id := primitive.NewObjectID()
	item := models.Item{
		ID:        id,
		Brand:     data.Brand,
		Model:     data.Model,
		Item_Name: data.Item_Name,
		Year:      data.Year,
		Price:     data.Price,
		CreatedAt: data.CreatedAt,
	}
	_, err := collection.InsertOne(context.Background(), item)

	return err
}
func (s *dbStore) GetAllItems(start int, end int, filters bson.D) ([]models.Item, int64, error) {
	// Calculate the limit and skip values for pagination
	limit := end - start + 1
	skip := start

	collection := s.db.Collection("item")
	// Set options for sorting or filtering if required
	findOptions := options.Find().SetSort(bson.M{"created_at": -1})
	// Apply pagination
	// 	skip := int64((pagination.PageNumber - 1) * pagination.PageSize)
	// 	findOptions.SetSkip(skip)
	// findOptions.SetLimit(int64(pagination.PageSize))

	// Query the database with limit and skip parameters
	findOptions.SetLimit(int64(limit)).SetSkip(int64(skip))
	// Find items with pagination
	cur, err := collection.Find(context.Background(), filters, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(context.Background())

	var itemslice []models.Item

	// Iterate over the result cursor
	for cur.Next(context.Background()) {
		var item models.Item
		if err := cur.Decode(&item); err != nil {
			return nil, 0, err
		}
		itemslice = append(itemslice, item)
	}

	if err := cur.Err(); err != nil {
		return nil, 0, err
	}
	if len(itemslice) == 0 {
		return nil, 0, errors.New("Documents not found")
	}
	totalCount, err := collection.CountDocuments(context.Background(), filters)
	if err != nil {
		return nil, 0, err
	}
	// w := context.Background().Value("responseWriter").(http.ResponseWriter)
	// Set the X-Total-Count header in the response
	// w.Header().Set("X-Total-Count", strconv.Itoa(int(totalCount)))
	return itemslice, totalCount, nil
}

func (s *dbStore) UpdateItemByID(filter bson.M, updater bson.M) error {
	collection := s.db.Collection("item")
	_, err := collection.UpdateOne(context.Background(), filter, updater)
	if err != nil {
		return errors.Wrap(err, "Errors in updating")
	}
	return nil
}
func (s *dbStore) UpdateItemByName(filter bson.M, updater bson.M) error {
	collection := s.db.Collection("item")
	_, err := collection.UpdateOne(context.Background(), filter, updater)
	if err != nil {
		return errors.Wrap(err, "Errors in updating")
	}
	return nil
}
func (s *dbStore) DeleteItemById(filter bson.M) error {
	collection := s.db.Collection("item")
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.Wrap(err, "Errors in deleting")
	}
	return nil
}

func (s *dbStore) DeleteItemByName(filter bson.M) error {
	collection := s.db.Collection("item")
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.Wrap(err, "Errors in deleting")
	}
	return nil
}

func (s *dbStore) Count(filters bson.D) (int64, error) {
	collection := s.db.Collection("item")
	return collection.CountDocuments(context.Background(), filters, &options.CountOptions{})
}
