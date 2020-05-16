package csv_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hsmtkk/post_crud/pkg/csv"
	"github.com/hsmtkk/post_crud/test/saveload"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "csv")
	assert.Nil(t, err, "should be nil")
	defer os.RemoveAll(tmpDir)
	csvPath := filepath.Join(tmpDir, "tmp.csv")
	assert.True(t, saveload.SaveAndLoad(csv.New(csvPath)), "should be true")
}
