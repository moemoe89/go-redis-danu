//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/practicing-redis-golang/api"

	"time"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	val := "val"
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	r := ap.NewRedisRepository(mock)

	err = r.Set(key, val, exp)
	assert.NoError(t, err)
}
/*
func TestSAdd(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	val := "val"

	mock := redismock.NewNiceMock(client)
	mock.On("SAdd", key, val).Return(redis.NewIntResult(0, nil))

	r := ap.NewRedisRepository(mock)

	err = r.SAdd(key, val)
	assert.NoError(t, err)
}
*/
func TestHSet(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	field := "field"
	val := "val"

	mock := redismock.NewNiceMock(client)
	mock.On("HSet", key, field, val).Return(redis.NewBoolResult(true, nil))

	r := ap.NewRedisRepository(mock)

	res, err := r.HSet(key, field, val)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestGet(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	val := "val"

	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := ap.NewRedisRepository(mock)
	res, err := r.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}
/*
func TestGetSAdd(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	val := []string{"val"}

	mock := redismock.NewNiceMock(client)
	mock.On("GetSadd", key).Return(redis.NewStringSliceResult(val, nil))

	r := ap.NewRedisRepository(mock)
	res, err := r.GetSAdd(key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}

func TestGetHSet(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"
	field := "field"
	val := "val"

	mock := redismock.NewNiceMock(client)
	mock.On("GetHSet", key, field).Return(redis.NewStringResult(val, nil))

	r := ap.NewRedisRepository(mock)
	res, err := r.GetHSet(key, field)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}
*/
func TestDelete(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	key := "key"

	mock := redismock.NewNiceMock(client)
	mock.On("Delete", key).Return(redis.NewStringResult("", nil))

	r := ap.NewRedisRepository(mock)
	err = r.Delete(key)
	assert.NoError(t, err)
}
