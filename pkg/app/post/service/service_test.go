package service_test

// import (
// 	appError "Project/internal/error"
// 	"Project/mocks"
// 	"Project/pkg/app/user/domain"
// 	"Project/pkg/app/user/service"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestCreate(t *testing.T) {
// 	testCases := []struct {
// 		Name      string
// 		setUp     func(MockUserStorage *mocks.MockUserStorageInterface) domain.CreateUserModel
// 		assertion func(err error, subTest *testing.T)
// 	}{
// 		{
// 			Name: "Succes",
// 			setUp: func(MockUserStorage *mocks.MockUserStorageInterface) domain.CreateUserModel {
// 				u := domain.CreateUserModel{
// 					Email:            "test",
// 					Password:         "test",
// 					Username:         "test",
// 					PasswordVerified: "test",
// 				}

// 				MockUserStorage.EXPECT().Create(gomock.Any()).Return(nil)

// 				return u
// 			},
// 			assertion: func(err error, subTest *testing.T) {
// 				assert.Nil(subTest, err)
// 			},
// 		},
// 		{
// 			Name: "Diferent password",
// 			setUp: func(MockUserStorage *mocks.MockUserStorageInterface) domain.CreateUserModel {
// 				u := domain.CreateUserModel{
// 					Email:            "test",
// 					Password:         "test",
// 					Username:         "test",
// 					PasswordVerified: "other",
// 				}
// 				// MockUserStorage.EXPECT().Create(gomock.Any()).Return(userError.DiferentPasswordError)
// 				return u

// 			},
// 			assertion: func(err error, subTest *testing.T) {
// 				assert.Equal(subTest, domain.DiferentPasswordError, err)
// 			},
// 		},
// 		{
// 			Name: "Error Database",
// 			setUp: func(MockUserStorage *mocks.MockUserStorageInterface) domain.CreateUserModel {
// 				u := domain.CreateUserModel{
// 					Email:            "test",
// 					Password:         "test",
// 					Username:         "test",
// 					PasswordVerified: "test",
// 				}
// 				MockUserStorage.EXPECT().Create(gomock.Any()).Return(errors.New("generic error"))
// 				return u
// 			},
// 			assertion: func(err error, subTest *testing.T) {
// 				app_error := &appError.AppError{}
// 				assert.False(subTest, errors.As(err, app_error))
// 			},
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	mockUserStorage := mocks.NewMockUserStorageInterface(ctrl)
// 	service := service.UserService{
// 		UserStorage: mockUserStorage,
// 	}
// 	for _, test := range testCases {

// 		t.Run(test.Name, func(subTest *testing.T) {

// 			u := test.setUp(mockUserStorage)

// 			err := service.CreateUserService(u)
// 			test.assertion(err, subTest)

// 		})
// 	}
// }
