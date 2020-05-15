package csv

import (
	"encoding/csv"
	"os"

	"github.com/hsmtkk/post_crud/pkg/post"
)

func New(csvPath string) post.PostStore {
	return &csvStore{csvPath: csvPath}
}

type csvStore struct {
	csvPath string
}

func (store *csvStore) Save(posts []post.Post) error {
	records := [][]string{}
	for _, p := range posts {
		records = append(records, p.Strings())
	}
	f, err := os.OpenFile(store.csvPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (store *csvStore) Load() ([]post.Post, error) {
	f, err := os.Open(store.csvPath)
	if err != nil {
		return nil, err
	}
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	posts := []post.Post{}
	for _, r := range records {
		p, err := post.NewFromStrings(r)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
