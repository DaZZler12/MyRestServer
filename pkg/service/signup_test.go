package service

import (
	"errors"
	"testing"

	mocks "github.com/DaZZler12/MyRestServer/mocks"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/serror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSignUp(t *testing.T) {
	type args struct {
		data models.User
	}

	testCases := []struct {
		name    string
		args    args
		srvc    *Service
		wantErr error
	}{
		{
			name: "db insert user fail",
			args: args{
				data: models.User{
					ID:       primitive.NewObjectID(),
					Name:     "test",
					Email:    "test@gmail.com",
					Country:  "India",
					Password: "Hello@123",
				},
			},
			srvc: func() *Service {
				getUserStore := new(mocks.Store)
				getUserStore.On("GetUserByEmail", mock.Anything).Return(models.User{}, errors.New("db get user error"))
				getUserStore.On("InsertUser", mock.Anything).Return(errors.New("db Insert not possible"))
				return &Service{
					store: getUserStore,
				}
			}(),
			wantErr: serror.BadRequestError("Failed to create an account"),
		},
		{
			name: "db user exists",
			args: args{
				data: models.User{
					ID:       primitive.NewObjectID(),
					Name:     "test",
					Email:    "test@gmail.com",
					Country:  "India",
					Password: "Hello@123",
				},
			},
			srvc: func() *Service {
				getUserStore := new(mocks.Store)
				getUserStore.On("GetUserByEmail", mock.Anything).Return(models.User{}, nil)
				return &Service{
					store: getUserStore,
				}
			}(),
			wantErr: serror.AlreadyInUse("Email is already in use"),
		},
		{
			name: "db insertion success",
			args: args{
				data: models.User{
					ID:       primitive.NewObjectID(),
					Name:     "test",
					Email:    "test@gmail.com",
					Country:  "India",
					Password: "Hello@123",
				},
			},
			srvc: func() *Service {
				getUserStore := new(mocks.Store)
				getUserStore.On("GetUserByEmail", mock.Anything).Return(models.User{}, errors.New("user does not exist"))
				getUserStore.On("InsertUser", mock.Anything).Return(nil)
				return &Service{
					store: getUserStore,
				}
			}(),
			wantErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.srvc.SignUp(tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
