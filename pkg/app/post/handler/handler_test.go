package handler_test

// import (
// 	"Project/mocks"
// 	"Project/pkg/app/user/domain"
// 	"Project/pkg/app/user/handler"
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestCreateUserHandler(t *testing.T) {
// 	testCases := []struct {
// 		Name      string
// 		setUp     func(MockUserService *mocks.MockUserServiceInterface) (*httptest.ResponseRecorder, *http.Request)
// 		assertion func(w *httptest.ResponseRecorder, subTest *testing.T)
// 	}{
// 		{
// 			Name: "Succes",
// 			setUp: func(MockUserService *mocks.MockUserServiceInterface) (*httptest.ResponseRecorder, *http.Request) {

// 				reqBody := `{
// 					"username": "steve",
// 					"email": "steve@gmail.com",
// 					"password": "clave123",
// 					"password_valid": "clave123"
// 				}
// 				`
// 				r, err := http.NewRequest(http.MethodPost, "user/",
// 					bytes.NewBuffer([]byte(reqBody)))

// 				if err != nil {
// 					t.Fatal("unexpected error creating request ", err.Error())
// 				}

// 				r.Header.Set("Content-Type", "application/json")

// 				w := httptest.NewRecorder()

// 				MockUserService.EXPECT().CreateUserService(gomock.Any()).Return(nil)
// 				return w, r
// 			},
// 			assertion: func(w *httptest.ResponseRecorder, subTest *testing.T) {
// 				assert.Equal(subTest, 201, w.Code)
// 			},
// 		},
// 		{
// 			Name: "Diferent password",
// 			setUp: func(MockUserService *mocks.MockUserServiceInterface) (*httptest.ResponseRecorder, *http.Request) {

// 				reqBody := `{
// 					"username": "steve",
// 					"email": "steve@gmail.com",
// 					"password": "clave123",
// 					"password_valid": "other password"
// 				}`

// 				r, err := http.NewRequest(http.MethodPost, "user/",
// 					bytes.NewBuffer([]byte(reqBody)))

// 				if err != nil {
// 					t.Fatal("unexpected error creating request ", err.Error())
// 				}

// 				r.Header.Set("Content-Type", "application/json")

// 				w := httptest.NewRecorder()

// 				MockUserService.EXPECT().CreateUserService(gomock.Any()).Return(domain.DiferentPasswordError)
// 				return w, r
// 			},
// 			assertion: func(w *httptest.ResponseRecorder, subTest *testing.T) {
// 				assert.Equal(subTest, 400, w.Code)
// 				assert.Equal(subTest, "{\"error\":{\"code\":\"password_diferents\",\"message\":\"the passwords are diferents\"}}\n", w.Body.String())
// 			},
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	mockUserService := mocks.NewMockUserServiceInterface(ctrl)
// 	handler := handler.UserHandler{
// 		UserService: mockUserService,
// 	}

// 	for _, test := range testCases {
// 		t.Run(test.Name, func(subTest *testing.T) {

// 			w, r := test.setUp(mockUserService)

// 			handler.CreateUserHandler(w, r)
// 			test.assertion(w, subTest)

// 		})
// 	}
// }
