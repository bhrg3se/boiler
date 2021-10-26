package store


import (
	"boiler/models"
	"crypto/rsa"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockStore) FetchUser(userID string) (*models.User, error) {
	args := m.Called(userID)
	r0, r1 := args.Get(0), args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*models.User), r1
}

func (m *MockStore) FetchUserWithPassword(email string) (*models.User, error) {
	args := m.Called(email)
	r0, r1 := args.Get(0), args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*models.User), r1
}

func (m *MockStore) GetConfig() models.Config {
	args := m.Called()
	return args.Get(0).(models.Config)
}

func (m *MockStore) GetJWTPublicKey() *rsa.PublicKey {
	args := m.Called()
	return args.Get(0).(*rsa.PublicKey)
}

func (m *MockStore) GetJWTPrivateKey() *rsa.PrivateKey {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey)
}



