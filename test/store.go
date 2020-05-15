package test

import "github.com/hsmtkk/post_crud/pkg/post"

func SaveAndLoad(store post.PostStore) bool {
	posts := []post.Post{}
	posts = append(posts, post.Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"})
	posts = append(posts, post.Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"})
	err := store.Save(posts)
	if err != nil {
		return false
	}
	loaded, err := store.Load()
	if err != nil {
		return false
	}
	if len(posts) != len(loaded) {
		return false
	}
	for i := range loaded {
		if posts[i] != loaded[i] {
			return false
		}
	}
	return true
}
