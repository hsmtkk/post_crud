package post

import (
	"encoding/json"
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

func (p *Post) ToJSON() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func FromJSON(js string) (Post, error) {
	p := Post{}
	if err := json.Unmarshal([]byte(js), &p); err != nil {
		return Post{}, err
	}
	return p, nil
}

type PostStore interface {
	Save([]Post) error
	Load() ([]Post, error)
}
