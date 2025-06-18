package boards

import (
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetFiles(t *testing.T) {
	o := Options{
		Folder: "testdata/basic",
	}
	files, err := o.getFiles(helpers.NewTestWriterContext(t))
	assert.NoError(t, err, "We shouldn't see an error yet")
	assert.Len(t, files, 2, "")
}

func TestGetFiles_InvalidFolder(t *testing.T) {
	o := Options{
		Folder: "does/not/exist",
	}
	files, err := o.getFiles(helpers.NewTestWriterContext(t))
	assert.Error(t, err, "We shouldn't see an error yet")
	assert.Nil(t, files)
}
