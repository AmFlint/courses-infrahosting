package datastore

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

const redisKey = "techtestgo_count"

type Datastore struct{
	client *redis.Client
}

// New - Create a new Datastore
func New(host, port string) (*Datastore, error) {
	connectionString := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr: connectionString,
		Password: "",
		DB: 0,
	})

	_, err := client.Get(redisKey).Int()
	if err != nil  {
		if err.Error() != "redis: nil" {
			return nil, err
		}

		if err := client.Set(redisKey, 0, 0).Err(); err != nil {
			log.Print(err)
			return nil, err
		}
	}

	return &Datastore{
		client: client,
	}, nil
}

func (d Datastore) GetCount() (int, error) {
	count := d.client.Get(redisKey)

	return count.Int()
}

func (d Datastore) SetCount(count int) error {
	return d.client.Set(redisKey, count, 0).Err()
}
