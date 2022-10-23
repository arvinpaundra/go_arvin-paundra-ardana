package controller

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/clean-architecture/database"
	"github.com/arvinpaundra/clean-architecture/dto"
	"github.com/arvinpaundra/clean-architecture/repository"
	"github.com/arvinpaundra/clean-architecture/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type suiteUsersController struct {
	suite.Suite
	mock               sqlmock.Sqlmock
	mockUserController *userController
}

func (s *suiteUsersController) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	database.DB = dbGorm
	s.mock = mock

	mockUserRepository := repository.NewUserRepository(database.DB)
	mockUserUsecase := usecase.NewUserUsecase(mockUserRepository)
	mockUserController := NewUserController(mockUserUsecase)

	s.mockUserController = mockUserController
}

func (s *suiteUsersController) TestHandlerLogin() {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ContentType        string
		Body               dto.UserDTO
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success login",
			ExpectedStatusCode: http.StatusOK,
			Method:             "POST",
			ContentType:        "application/json",
			Body: dto.UserDTO{
				Email:    "arvin@mail.com",
				Password: "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "login success",
				"data": map[string]interface{}{
					"user": map[string]interface{}{
						"email":    "arvin@mail.com",
						"password": "123",
					},
					//"token": "",
				},
			},
		},
		{
			Name:               "failed without content type",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             "POST",
			ContentType:        "",
			Body:               dto.UserDTO{},
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "error",
				"message": "code=415, message=Unsupported Media Type",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL LIMIT 1")).
				WithArgs(v.Body.Email, v.Body.Password).
				WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).AddRow(v.Body.Email, v.Body.Password))

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/login", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")
			ctx.Request().Header.Set("Content-Type", v.ContentType)

			err := s.mockUserController.HandlerLogin(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult, resp)
			}
		})
	}
}

func (s *suiteUsersController) TestHandlerGetAllUsers() {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get all users",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all users",
				"data": map[string]interface{}{
					"users": []interface{}{
						map[string]interface{}{
							"email":    "arvin@mail.com",
							"password": "123",
						},
					},
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).AddRow("arvin@mail.com", "123"))

			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := s.mockUserController.HandlerGetAllUsers(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult, resp)
			}
		})
	}
}

func (s *suiteUsersController) TestHandlerCreateUser() {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ContentType        string
		Body               dto.UserDTO
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success create user",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			ContentType:        "application/json",
			Body: dto.UserDTO{
				Email:    "arvin@mail.com",
				Password: "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success create user",
				"data": map[string]interface{}{
					"user": map[string]interface{}{
						"email":    "arvin@mail.com",
						"password": "123",
					},
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).
				WithArgs(Anytime{}, Anytime{}, nil, v.Body.Email, v.Body.Password).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")
			ctx.Request().Header.Set("Content-Type", v.ContentType)

			err := s.mockUserController.HandlerCreateUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult, resp)
			}
		})
	}
}

func (s *suiteUsersController) TearDownSuite() {
	database.DB = nil
	s.mock = nil
}

func TestSuiteUserController(t *testing.T) {
	suite.Run(t, new(suiteUsersController))
}
