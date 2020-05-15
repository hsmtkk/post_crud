package post

import (
	"fmt"
	"strconv"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func (p *Post) Strings() []string {
	return []string{fmt.Sprintf("%d", p.ID), p.Content, p.Author}
}

func NewFromStrings(ss []string) (Post, error) {
	if len(ss) != 3 {
		return Post{}, fmt.Errorf("invalid length")
	}
	id, err := strconv.Atoi(ss[0])
	if err != nil {
		return Post{}, err
	}
	return Post{ID: id, Content: ss[1], Author: ss[2]}, nil
}

type PostStore interface {
	Save([]Post) error
	Load() ([]Post, error)
}
