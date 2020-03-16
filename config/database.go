//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"fmt"

	"github.com/go-redis/redis"
)

// InitDB will create a variable that represent the redis.Client
func InitDB() (*redis.Client, error) {
	Client := redis.NewClient(&redis.Options{
		Addr:     Configuration.Redis.Addr,
		Password: Configuration.Redis.Password,
		DB:       Configuration.Redis.DB,
	})

	res := Client.Ping()
	if res.Err() != nil {
		return nil, fmt.Errorf("Failed to ping connection to redis: %s", res.Err().Error())
	}

	return Client, nil
}
