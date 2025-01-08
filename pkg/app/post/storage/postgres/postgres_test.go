package postgres_test

// import (
// 	"Project/mocks"
// 	"Project/pkg/app/user/domain"
// 	Storage "Project/pkg/app/user/storage/postgres"
// 	"testing"
// 	"time"

// 	"github.com/lib/pq"

// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestCreate(t *testing.T) {
// 	testCases := []struct {
// 		Name      string
// 		setUp     func(MockUserStorage *mocks.MockDBinterface) domain.CreateUserModel
// 		assertion func(err error, subTest *testing.T)
// 	}{
// 		{
// 			Name: "Success",
// 			setUp: func(MockUserStorage *mocks.MockDBinterface) domain.CreateUserModel {
// 				u := domain.CreateUserModel{
// 					Email:            "test",
// 					Password:         "test",
// 					Username:         "test",
// 					PasswordVerified: "test",
// 				}
// 				u.CreatedAt = time.Now().UTC()

// 				MockUserStorage.EXPECT().Exec("insert into users (username, email, password, created_at) values ($1,$2,$3,$4)",
// 					u.Username, u.Email, u.Password, u.CreatedAt).Return(nil, nil)

// 				return u
// 			},
// 			assertion: func(err error, subTest *testing.T) {
// 				assert.Nil(subTest, err)
// 			},
// 		},
// 		{
// 			Name: "Email already exist",
// 			setUp: func(MockUserStorage *mocks.MockDBinterface) domain.CreateUserModel {
// 				u := domain.CreateUserModel{
// 					Email:            "test",
// 					Password:         "test",
// 					Username:         "test",
// 					PasswordVerified: "test",
// 				}
// 				u.CreatedAt = time.Now().UTC()
// 				pgErr := &pq.Error{
// 					Code:    "23505",
// 					Message: "duplicate key value violates unique constraint",
// 				}

// 				var err error = pgErr
// 				MockUserStorage.EXPECT().Exec("insert into users (username, email, password, created_at) values ($1,$2,$3,$4)",
// 					u.Username, u.Email, u.Password, u.CreatedAt).Return(nil, err)

// 				return u
// 			},
// 			assertion: func(err error, subTest *testing.T) {
// 				assert.Equal(subTest, domain.EmailAlreadyExistError, err)
// 			},
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	mockDB := mocks.NewMockDBinterface(ctrl)

// 	storage := Storage.UserStorage{DB: mockDB}

// 	for _, test := range testCases {
// 		t.Run(test.Name, func(subTest *testing.T) {

// 			u := test.setUp(mockDB)

// 			err := storage.Create(u)
// 			test.assertion(err, subTest)

// 		})
// 	}
// }
