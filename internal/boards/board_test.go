package boards

import (
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	o := Options{
		Folder: "testdata/basic",
	}
	response, err := o.Combine(helpers.NewTestWriterContext(t))
	assert.Nil(t, err)
	assert.Equal(t, "A1-100X", response.Boards[0].Name)
	assert.Equal(t, "B7-400X", response.Boards[1].Name)
	assert.Equal(t, "C1-100X", response.Boards[2].Name)
	assert.Equal(t, "D4-200S", response.Boards[3].Name)
	assert.Equal(t, "Low_Power", response.Boards[4].Name)
	assert.Equal(t, 5, response.Metadata.Totals.Boards)
	assert.Equal(t, 2, response.Metadata.Totals.Vendors)
	assert.Equal(t, false, response.Metadata.Errors.HasErrors)
}

func TestCombine_WithInvalidJsonFilesInFolders(t *testing.T) {
	o := Options{
		Folder: "testdata/invalidMix",
	}
	response, err := o.Combine(helpers.NewTestWriterContext(t))
	assert.Nil(t, err)
	assert.Equal(t, 3, response.Metadata.Totals.Boards)
	assert.Equal(t, 2, response.Metadata.Totals.Vendors)
	assert.Equal(t, true, response.Metadata.Errors.HasErrors)
	assert.Equal(t, 1, response.Metadata.Errors.FileReadErrors)
	assert.Contains(t, response.Metadata.Errors.Files, "testdata/invalidMix/invalid.json")
}

func TestCombine_InvalidFolder(t *testing.T) {
	o := Options{
		Folder: "where/the/heck/am/i",
	}
	_, err := o.Combine(helpers.NewTestWriterContext(t))
	assert.NotNil(t, err)
}
