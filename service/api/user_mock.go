package api

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context) error {

	args := m.Called()

	return args.Error(1)

}
