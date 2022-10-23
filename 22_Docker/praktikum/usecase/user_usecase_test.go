package usecase

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/clean-architecture/database"
	"github.com/arvinpaundra/clean-architecture/repository"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type suiteUserUsecase struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	mockUserUsecase *userUsecase
}

func (s *suiteUserUsecase) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	database.DB = dbGorm
	s.mock = mock

	userRepository := repository.NewUserRepository(database.DB)
	s.mockUserUsecase = NewUserUsecase(userRepository)
}

func (s *suiteUserUsecase) TestLogin() {

}

func (s *suiteUserUsecase) TearDownSuite() {
	database.DB = nil
	s.mock = nil
}

func TestSuiteUserUsecase(t *testing.T) {
	suite.Run(t, new(suiteUserUsecase))
}
