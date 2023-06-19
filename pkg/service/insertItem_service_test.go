package service

import (
	"errors"
	"testing"

	mocks "github.com/DaZZler12/MyRestServer/mocks"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertItem(t *testing.T) {
	type args struct {
		data models.Item
	}

	testCases := []struct {
		name    string
		args    args
		srvc    *Service
		wantErr error
	}{
		{
			name: "item already exists",
			args: args{
				data: models.Item{
					ID:        primitive.NewObjectID(),
					Brand:     "Brand A",
					Model:     "Model A",
					Item_Name: "Item A",
					Year:      2021,
					Price:     9.99,
				},
			},
			srvc: func() *Service {
				store := new(mocks.Store)
				filter := bson.M{"item_name": "Item A"}
				store.On("GetItemByModelBrand", filter).Return(models.Item{}, nil)
				return &Service{
					store: store,
				}
			}(),
			wantErr: serror.AlreadyInUse("Item Details Already there"),
		},
		{
			name: "insertion success",
			args: args{
				data: models.Item{
					ID:        primitive.NewObjectID(),
					Brand:     "Brand B",
					Model:     "Model B",
					Item_Name: "Item B",
					Year:      2022,
					Price:     19.99,
				},
			},
			srvc: func() *Service {
				store := new(mocks.Store)
				filter := bson.M{"item_name": "Item B"}
				store.On("GetItemByModelBrand", filter).Return(models.Item{}, serror.NotFoundError("Item not found"))
				store.On("InsertItem", mock.Anything).Return(nil)
				return &Service{
					store: store,
				}
			}(),
			wantErr: nil,
		},
		{
			name: "insertion failure",
			args: args{
				data: models.Item{
					ID:        primitive.NewObjectID(),
					Brand:     "Brand C",
					Model:     "Model C",
					Item_Name: "Item C",
					Year:      2023,
					Price:     29.99,
				},
			},
			srvc: func() *Service {
				store := new(mocks.Store)
				filter := bson.M{"item_name": "Item C"}
				store.On("GetItemByModelBrand", filter).Return(models.Item{}, serror.NotFoundError("Item not found"))
				store.On("InsertItem", mock.Anything).Return(errors.New("database error"))
				return &Service{
					store: store,
				}
			}(),
			wantErr: serror.BadRequestError("Error adding the Item"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.srvc.InsertItem(tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
