//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/practicing-redis-golang/api"
	"github.com/moemoe89/practicing-redis-golang/api/mocks"
	"github.com/moemoe89/practicing-redis-golang/config"

	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceSet(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	key := "key"
	value := "value"
	duration := time.Duration(0)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Set", key, value, duration).Return(nil).Once()
		s := ap.NewService(log, mockRepo)

		status, err := s.Set(key, value, duration)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Set", key, value, duration).Return(errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		status, err := s.Set(key, value, duration)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceHSet(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	key := "key"
	field := "field"
	value := "value"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("HSet", key, field, value).Return(true, nil).Once()
		s := ap.NewService(log, mockRepo)

		isCreate, status, err := s.HSet(key, field, value)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, true, isCreate)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("HSet", key, field, value).Return(false, errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		isCreate, status, err := s.HSet(key, field, value)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, false, isCreate)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceGet(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	key := "key"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Get", key).Return("value", nil).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.Get(key)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, "value", val)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Get", key).Return("", errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.Get(key)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, "", val)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceGetSAdd(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	key := "key"
	empArr := []string{}
	notEmpArr := []string{"value"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetSAdd", key).Return(notEmpArr, nil).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.GetSAdd(key)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, notEmpArr, val)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("GetSAdd", key).Return(empArr, errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.GetSAdd(key)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
		assert.Empty(t, val)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceGetHSet(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	key := "key"
	field := "field"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetHSet", key, field).Return("value", nil).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.GetHSet(key, field)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, "value", val)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("GetHSet", key, field).Return("", errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		val, status, err := s.GetHSet(key, field)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, "", val)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceDelete(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		s := ap.NewService(log, mockRepo)

		status, err := s.Delete("key")

		assert.NoError(t, err)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("Unexpected database error")).Once()
		s := ap.NewService(log, mockRepo)

		status, err := s.Delete("key")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}