package main

import (
	"log"

	"github.com/hsmtkk/post_crud/pkg/redis"
	"github.com/hsmtkk/post_crud/test/saveload"
)

func main() {
	store, err := redis.New("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	success := saveload.SaveAndLoad(store)
	if !success {
		log.Fatal("test failed")
	}
}
