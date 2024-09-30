package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddition(t *testing.T) {
	result := 2 + 2

	assert.Equal(t, 4, result, "The result should be 4")
}

// wrappers!
func TestAssertions(t *testing.T) {
	assert.Equal(t, 1, 1, "1 should be equal to 1")
	assert.NotEqual(t, 1, 2, "1 should not be equal to 2")
	assert.Nil(t, nil, "Nil should be equal to nil")
	assert.NotNil(t, "test", "The string should not be nil")
	assert.True(t, true, "True should be true")
	assert.False(t, false, "False should be false")
	assert.Empty(t, "", "The string should be empty")
	assert.NotEmpty(t, "test", "The string should not be empty")
}

// -----------------------------------------------------------------------

type DB interface {
	Get(id string) (string, error)
	Put(id string, value string) error
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Get(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockDB) Put(id string, value string) error {
	args := m.Called(id, value)
	return args.Error(0)
}

func TestMocking(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("Get", "test-id").Return("test-value", nil)
	mockDB.On("Put", "test-id", "test-value").Return(nil)

	// ***

	value, err := mockDB.Get("test-id")
	assert.NoError(t, err)
	assert.Equal(t, "test-value", value)

	err = mockDB.Put("test-id", "test-value")
	assert.NoError(t, err)

	// ***

	mockDB.AssertExpectations(t)
}
