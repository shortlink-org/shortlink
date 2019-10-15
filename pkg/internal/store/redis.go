package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/internal/link"
	"github.com/go-redis/redis"
)

type RedisLinkList struct {
	client *redis.Client
}

func (r *RedisLinkList) Init() error {
	// Connect to Redis
	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := r.client.Ping().Result(); err != nil {
		return err
	}

	return nil
}

func (r *RedisLinkList) Get(id string) (*link.Link, error) {
	val, err := r.client.Get(id).Result()
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response link.Link

	if err = json.Unmarshal([]byte(val), &response); err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed parse link: %s", id))}
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response, nil
}

func (r *RedisLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	val, err := json.Marshal(data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: errors.New(fmt.Sprintf("Failed marsharing link: %s", data.Url))}
	}

	err = r.client.Set(data.Hash, val, 0).Err()
	if err = r.client.Set(data.Hash, val, 0).Err(); err != nil {
		return nil, &link.NotFoundError{Link: data, Err: errors.New(fmt.Sprintf("Failed save link: %s", data.Url))}
	}

	return &data, nil
}

func (r *RedisLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (r *RedisLinkList) Delete(id string) error {
	if err := r.client.Del(id).Err(); err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed save link: %s", id))}
	}

	return nil
}
