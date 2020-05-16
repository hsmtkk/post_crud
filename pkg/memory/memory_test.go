package memory_test

import (
	"testing"

	"github.com/hsmtkk/post_crud/pkg/memory"
	"github.com/hsmtkk/post_crud/test/saveload"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert.True(t, saveload.SaveAndLoad(memory.New()), "should be true")
}
