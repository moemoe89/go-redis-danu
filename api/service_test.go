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
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceDelete(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete("key")

		assert.NoError(t, err)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-delete", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete("key")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

}