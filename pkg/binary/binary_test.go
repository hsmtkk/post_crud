package binary_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hsmtkk/post_crud/pkg/binary"
	"github.com/hsmtkk/post_crud/test"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "binary")
	assert.Nil(t, err, "should be nil")
	defer os.RemoveAll(tmpDir)
	binaryPath := filepath.Join(tmpDir, "tmp.bin")
	assert.True(t, test.SaveAndLoad(binary.New(binaryPath)), "should be true")
}
