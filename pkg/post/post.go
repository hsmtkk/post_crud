package post

type Post struct {
	ID      int
	Content string
	Author  string
}

type PostStore interface {
	Save([]Post) error
	Load() ([]Post, error)
}
