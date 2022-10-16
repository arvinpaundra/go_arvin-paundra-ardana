package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/unit-test/config"
	"github.com/arvinpaundra/unit-test/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type suiteUsers struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func (s *suiteUsers) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	config.DB = dbGorm
	s.mock = mock
}

func (s *suiteUsers) TestLoginUserController() {
	row := sqlmock.NewRows([]string{"name", "email", "password"}).AddRow("Arvin Paundra", "arvin@mail.com", "123")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               DTOUser
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success login",
			ExpectedStatusCode: http.StatusOK,
			Method:             "POST",
			Body: DTOUser{
				Name:     "Arvin Paundra",
				Email:    "arvin@mail.com",
				Password: "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"user": map[string]interface{}{
					"name":     "Arvin Paundra",
					"email":    "arvin@mail.com",
					"password": "123",
				},
				"token": "",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/login", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := LoginUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["user"], resp["user"])
			}
		})
	}
}

func (s *suiteUsers) TestGetUsersController() {
	row := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow("Arvin", "arvin@mail.com", "1234")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`name`,`users`.`email`,`users`.`password` FROM `users` WHERE `users`.`deleted_at` IS NULL")).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               []models.User
		HasReturnBody      bool
		ExpectedBody       []models.User
	}{
		{
			Name:               "success get users",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			Body: []models.User{
				{
					Name:     "Arvin",
					Email:    "arvin@mail.com",
					Password: "1234",
				},
			},
			HasReturnBody: true,
			ExpectedBody: []models.User{
				{
					Name:     "Arvin",
					Email:    "arvin@mail.com",
					Password: "1234",
				},
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := GetUsersController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string][]models.User
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["users"])
			}
		})
	}
}

func (s *suiteUsers) TestGetUserController() {
	row := sqlmock.NewRows([]string{"name", "email", "password"}).AddRow("Arvin", "arvin@mail.com", "1234")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`name`,`users`.`email`,`users`.`password` FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT 1")).
		WithArgs(1).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.User
		HasReturnBody      bool
		ExpectedBody       models.User
	}{
		{
			Name:               "success get user",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			Body: models.User{
				Name:     "Arvin",
				Email:    "arvin@mail.com",
				Password: "1234",
			},
			HasReturnBody: true,
			ExpectedBody: models.User{
				Name:     "Arvin",
				Email:    "arvin@mail.com",
				Password: "1234",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]models.User
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["user"])
			}
		})
	}
}

func (s *suiteUsers) TestCreateUserController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`) VALUES (?,?,?,?,?,?)")).
		WithArgs(Anytime{}, Anytime{}, nil, "Arvin", "arvin@mail.com", "123").
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)
	s.mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.User
		HasReturnBody      bool
		ExpectedBody       models.User
	}{
		{
			Name:               "success create user",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Body: models.User{
				Name:     "Arvin",
				Email:    "arvin@mail.com",
				Password: "123",
			},
			HasReturnBody: true,
			ExpectedBody: models.User{
				Name:     "Arvin",
				Email:    "arvin@mail.com",
				Password: "123",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := CreateUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]models.User
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["user"])
			}
		})
	}
}

func (s *suiteUsers) TestUpdateUserController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET")).
		WithArgs(Anytime{}, "Updated Name", "Updated Email", "Updated Password", 1).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)
	s.mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               DTOUser
		HasReturnBody      bool
		ExpectedBody       DTOUser
	}{
		{
			Name:               "success update user",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Body: DTOUser{
				Name:     "Updated Name",
				Email:    "Updated Email",
				Password: "Updated Password",
			},
			HasReturnBody: true,
			ExpectedBody: DTOUser{
				Name:     "Updated Name",
				Email:    "Updated Email",
				Password: "Updated Password",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := UpdateUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]DTOUser
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["user"])
			}
		})
	}
}

func (s *suiteUsers) TestDeleteUserController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)
	s.mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			Name:               "success delete",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedBody:       "success delete user",
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := DeleteUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]string
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["message"])
			}
		})
	}
}

func (s *suiteUsers) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
