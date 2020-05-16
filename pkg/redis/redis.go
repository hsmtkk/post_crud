package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/hsmtkk/post_crud/pkg/post"
)

func New(redisHost string) (post.PostStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &redisStore{client: client}, nil
}

type redisStore struct {
	client *redis.Client
}

func (store *redisStore) Save(posts []post.Post) error {
	for _, p := range posts {
		js, err := p.ToJSON()
		if err != nil {
			return err
		}
		if status := store.client.Set(string(p.ID), js, 0); status != nil {
			return fmt.Errorf("failed to set key-value to redis")
		}
	}
	return nil
}

func (store *redisStore) Load() ([]post.Post, error) {
	posts := []post.Post{}
	keys, err := store.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		val, err := store.client.Get(key).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to get key-value to redis")
		}
		p, err := post.FromJSON(val)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
