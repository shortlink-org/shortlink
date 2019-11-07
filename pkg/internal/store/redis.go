package store

import (
	"encoding/json"
	"fmt"

	"github.com/batazor/shortlink/pkg/link"
	"github.com/go-redis/redis"
)

// RedisLinkList implementation of store interface
type RedisLinkList struct { // nolint unused
	client *redis.Client
}

// Init ...
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

// Get ...
func (r *RedisLinkList) Get(id string) (*link.Link, error) {
	val, err := r.client.Get(id).Result()
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	if err = json.Unmarshal([]byte(val), &response); err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	if response.URL == "" {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (r *RedisLinkList) List() ([]*link.Link, error) {
	panic("implement me")
}

// Add ...
func (r *RedisLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	val, err := json.Marshal(data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed marsharing link: %s", data.URL)}
	}

	if err = r.client.Set(data.Hash, val, 0).Err(); err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.URL)}
	}

	return &data, nil
}

// Update ...
func (r *RedisLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (r *RedisLinkList) Delete(id string) error {
	if err := r.client.Del(id).Err(); err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}
