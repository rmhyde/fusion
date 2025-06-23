package boards

import (
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetFilesRecursively(t *testing.T) {
	o := Options{
		Ctx:       helpers.NewTestWriterContext(t),
		Folder:    "testdata/basic",
		Recursive: false,
	}
	files, err := o.getFilesRecursively()
	assert.NoError(t, err, "We shouldn't see an error")
	assert.Len(t, files, 2)

	o.Recursive = true
	files, err = o.getFilesRecursively()
	assert.NoError(t, err, "We still shouldn't see an error")
	assert.Len(t, files, 4)
}

func TestGetFilesRecursively_InvalidFolder(t *testing.T) {
	o := Options{
		Ctx:    helpers.NewTestWriterContext(t),
		Folder: "does/not/exist",
	}
	files, err := o.getFilesRecursively()
	assert.Error(t, err, "We should see an error")
	assert.Nil(t, files)
}
