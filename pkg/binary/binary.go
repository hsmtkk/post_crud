package binary

import (
	"bytes"
	"encoding/gob"
	"io"
	"io/ioutil"

	"github.com/hsmtkk/post_crud/pkg/post"
)

func New(binaryPath string) post.PostStore {
	return &binaryStore{binaryPath: binaryPath}
}

type binaryStore struct {
	binaryPath string
}

func (store *binaryStore) Save(posts []post.Post) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	for _, p := range posts {
		err := encoder.Encode(&p)
		if err != nil {
			return err
		}
	}
	err := ioutil.WriteFile(store.binaryPath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (store *binaryStore) Load() ([]post.Post, error) {
	bs, err := ioutil.ReadFile(store.binaryPath)
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	posts := []post.Post{}
	for {
		p := post.Post{}
		err := decoder.Decode(&p)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
