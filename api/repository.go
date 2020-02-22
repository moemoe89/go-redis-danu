//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"time"

	"gopkg.in/redis.v5"
)

// Repository represent the repositories
type Repository interface {
	Set(key string, value interface{}, exp time.Duration) error
	SAdd(key string, value ...interface{}) error
	HSet(key, field, value string) (bool, error)
	Get(key string) (string, error)
	GetHSet(key, field string) (string, error)
	GetSadd(key string) ([]string, error)
	Delete(key string) error
}

type redisRepository struct {
	Client *redis.Client
}

// NewRedisRepository will create an object that represent the Repository interface
func NewRedisRepository(Client *redis.Client) Repository {
	return &redisRepository{Client}
}

func (r *redisRepository) Set(key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(key, value, exp).Err()
}

func (r *redisRepository) SAdd(key string, value ...interface{}) error {
	return r.Client.SAdd(key, value).Err()
}

func (r *redisRepository) HSet(key, field, value string) (bool, error) {
	hset := r.Client.HSet(key, field, value)
	return hset.Val(), hset.Err()
}

func (r *redisRepository) Get(key  string) (string, error) {
	get := r.Client.Get(key)
	return get.Result()
}

func (r *redisRepository) GetSadd(key  string) ([]string, error) {
	get := r.Client.SMembers(key)
	return get.Result()
}

func (r *redisRepository) GetHSet(key, field  string) (string, error) {
	get := r.Client.HGet(key,field)
	return get.Result()
}

func (r *redisRepository) Delete(key string) error {
	return r.Client.Del(key).Err()
}