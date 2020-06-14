//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-redis-danu/config"
	"github.com/moemoe89/go-redis-danu/routers"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	log := config.InitLog()

	router := routers.GetRouter(log, nil)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
