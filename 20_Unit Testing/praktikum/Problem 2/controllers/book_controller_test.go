package controllers

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/unit-test/config"
	"github.com/arvinpaundra/unit-test/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type suiteBooks struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func (s *suiteBooks) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	config.DB = dbGorm
	s.mock = mock
}

func (s *suiteBooks) TestGetBooksController() {
	row := sqlmock.NewRows([]string{"title", "author", "publisher"}).AddRow("Title Test 1", "Author Test 1", "Publisher Test 1")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `books`.`title`,`books`.`author`,`books`.`publisher` FROM `books` WHERE `books`.`deleted_at` IS NULL")).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               []models.Book
		HasReturnBody      bool
		ExpectedBody       []models.Book
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			Body: []models.Book{
				{
					Title:     "Title Test 1",
					Author:    "Author Test 1",
					Publisher: "Publisher Test 1",
				},
			},
			HasReturnBody: true,
			ExpectedBody: []models.Book{
				{
					Title:     "Title Test 1",
					Author:    "Author Test 1",
					Publisher: "Publisher Test 1",
				},
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/books", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := GetBooksController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string][]models.Book
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["books"])
			}
		})
	}
}

func (s *suiteBooks) TestGetBookController() {
	row := sqlmock.NewRows([]string{"title", "author", "publisher"}).AddRow("Title Test 1", "Author Test 1", "Publisher Test 1")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `books`.`title`,`books`.`author`,`books`.`publisher` FROM `books` WHERE id = ? AND `books`.`deleted_at` IS NULL LIMIT 1")).
		WithArgs(1).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.Book
		HasReturnBody      bool
		ExpectedBody       models.Book
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			Body: models.Book{
				Title:     "Title Test 1",
				Author:    "Author Test 1",
				Publisher: "Publisher Test 1",
			},
			HasReturnBody: true,
			ExpectedBody: models.Book{
				Title:     "Title Test 1",
				Author:    "Author Test 1",
				Publisher: "Publisher Test 1",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/books", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetBookController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]models.Book
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["book"])
			}
		})
	}
}

func (s *suiteBooks) TestCreateBookController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`title`,`author`,`publisher`) VALUES (?,?,?,?,?,?)")).
		WithArgs(Anytime{}, Anytime{}, nil, "Title Test 1", "Author Test 1", "Publisher Test 1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               DTOBook
		HasReturnBody      bool
		ExpectedBody       DTOBook
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Body: DTOBook{
				Title:     "Title Test 1",
				Author:    "Author Test 1",
				Publisher: "Publisher Test 1",
			},
			HasReturnBody: true,
			ExpectedBody: DTOBook{
				Title:     "Title Test 1",
				Author:    "Author Test 1",
				Publisher: "Publisher Test 1",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/books", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := CreateBookController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]DTOBook
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["book"])
			}
		})
	}
}

func (s *suiteBooks) TestUpdateBookController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE `books` SET").
		WithArgs(Anytime{}, "Updated Title", "Updated Author", "Updated Publisher", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               DTOBook
		HasReturnBody      bool
		ExpectedBody       DTOBook
	}{
		{
			Name:               "success update book",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Body: DTOBook{
				Title:     "Updated Title",
				Author:    "Updated Author",
				Publisher: "Updated Publisher",
			},
			HasReturnBody: true,
			ExpectedBody: DTOBook{
				Title:     "Updated Title",
				Author:    "Updated Author",
				Publisher: "Updated Publisher",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/books", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := UpdateBookController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]DTOBook
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody, resp["book"])
			}
		})
	}
}

func (s *suiteBooks) TestDeleteBookController() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `books` WHERE id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 0)).
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
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedBody:       "success delete book",
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/books", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := DeleteBookController(ctx)
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

func (s *suiteBooks) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

func TestSuiteBooks(t *testing.T) {
	suite.Run(t, new(suiteBooks))
}
