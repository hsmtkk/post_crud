package memory

import (
	"github.com/hsmtkk/post_crud/pkg/post"
)

func New() post.PostStore {
	return &memoryStore{}
}

type memoryStore struct {
	posts []post.Post
}

func (store *memoryStore) Save(posts []post.Post) error {
	store.posts = []post.Post{}
	for _, p := range posts {
		store.posts = append(store.posts, p)
	}
	return nil
}

func (store *memoryStore) Load() ([]post.Post, error) {
	ps := []post.Post{}
	for _, p := range store.posts {
		ps = append(ps, p)
	}
	return ps, nil
}
