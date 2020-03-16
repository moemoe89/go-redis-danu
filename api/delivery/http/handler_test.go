//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package http_test

import (
	"github.com/moemoe89/practicing-redis-golang/api/mocks"
	"github.com/moemoe89/practicing-redis-golang/config"
	"github.com/moemoe89/practicing-redis-golang/routers"

	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteSuccess(t *testing.T) {
	log := config.InitLog()
	key := "key"

	mockService := new(mocks.Service)
	mockService.On("Delete", key).Return(http.StatusOK, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/delete/"+key, nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeleteFail(t *testing.T) {
	log := config.InitLog()
	key := "key"

	mockService := new(mocks.Service)
	mockService.On("Delete", key).Return(http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/delete/"+key, nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}