package boards

import (
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	o := Options{
		Folder: "testdata/basic",
		Ctx:    helpers.NewTestWriterContext(t),
	}
	response, err := o.Combine()
	assert.Nil(t, err)
	assert.Equal(t, "A1-100X", response.Boards[0].Name)
	assert.Equal(t, "B7-400X", response.Boards[1].Name)
	assert.Equal(t, "C1-100X", response.Boards[2].Name)
	assert.Equal(t, "D4-200S", response.Boards[3].Name)
	assert.Equal(t, "Shrug Board", response.Boards[4].Name)
	assert.Equal(t, "Low_Power", response.Boards[5].Name)
	assert.Equal(t, 6, response.Metadata.Totals.Boards)
	assert.Equal(t, 3, response.Metadata.Totals.Vendors)
	assert.Equal(t, 3, response.Metadata.Totals.WifiEnabled)
	assert.Equal(t, false, response.Metadata.Errors.HasErrors)
}

func TestCombine_WithInvalidJsonFilesInFolders(t *testing.T) {
	o := Options{
		Folder: "testdata/invalidMix",
		Ctx:    helpers.NewTestWriterContext(t),
	}
	response, err := o.Combine()
	assert.Nil(t, err)
	assert.Equal(t, 3, response.Metadata.Totals.Boards)
	assert.Equal(t, 2, response.Metadata.Totals.Vendors)
	assert.Equal(t, true, response.Metadata.Errors.HasErrors)
	assert.Equal(t, 2, response.Metadata.Errors.FileReadErrors)
	assert.Contains(t, response.Metadata.Errors.Files, "testdata/invalidMix/boards-broken.json")
	assert.Contains(t, response.Metadata.Errors.Files, "testdata/invalidMix/invalid.json")
}

func TestCombine_InvalidFolder(t *testing.T) {
	o := Options{
		Folder: "where/the/heck/am/i",
		Ctx:    helpers.NewTestWriterContext(t),
	}
	_, err := o.Combine()
	assert.NotNil(t, err)
}
