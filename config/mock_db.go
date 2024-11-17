// config/mock_db.go
package config

import (
	"jessie_miniproject/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// DBInterface adalah interface untuk menggantikan *gorm.DB dalam pengujian
type DBInterface interface {
	Find(out interface{}, where ...interface{}) *gorm.DB
}

// MockDB adalah mock dari DBInterface
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) SetFindResult(result interface{}, err error) {
	m.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]models.ProductLog)
		*arg = result.([]models.ProductLog)
	}).Return(&gorm.DB{Error: err})
}
