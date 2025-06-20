package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name            string
	TestingGenerics bool `json:"testing_generics"`
}

func TestReadAsType_FileDoesNotExistGiveError(t *testing.T) {
	_, err := ReadAsType[TestStruct]("there_is_no.cake")
	assert.Error(t, err)
}

func TestReadAsType_FailureToUnmarshalIsHandled(t *testing.T) {
	_, err := ReadAsType[TestStruct]("testdata/badJsonTest.json")
	assert.Error(t, err)
}

func TestReadAsType_FileCanBeRead(t *testing.T) {
	contents, err := ReadAsType[TestStruct]("testdata/genericReadTest.json")
	assert.Nil(t, err)
	assert.Equal(t, contents.Name, "example name")
	assert.Equal(t, contents.TestingGenerics, true)
}
