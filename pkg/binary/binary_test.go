package binary_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hsmtkk/post_crud/pkg/binary"
	"github.com/hsmtkk/post_crud/test/saveload"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "binary")
	assert.Nil(t, err, "should be nil")
	defer os.RemoveAll(tmpDir)
	binaryPath := filepath.Join(tmpDir, "tmp.bin")
	assert.True(t, saveload.SaveAndLoad(binary.New(binaryPath)), "should be true")
}
