package models

import (
	"gopkg.in/redis.v5"
)

var client *redis.Client

func init(){

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}