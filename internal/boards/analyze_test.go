package boards

import (
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSortAndGatherMetrics(t *testing.T) {
	wrapper := inputData()
	o := Options{}
	response := o.sortAndGatherMetrics(helpers.NewTestWriterContext(t), wrapper)
	assert.Equal(t, expectedData(), response, "Lists do not match, something has gone wrong")
}

func inputData() BoardWrapper {
	return BoardWrapper{
		Boards: []Board{
			{
				Name:    "Z1",
				Vendor:  "V1",
				Core:    "C1",
				HasWifi: false,
			},
			{
				Name:    "A1",
				Vendor:  "v1",
				Core:    "C1",
				HasWifi: false,
			},
			{
				Name:    "D1",
				Vendor:  "A1",
				Core:    "C1",
				HasWifi: true,
			},
		},
		Metadata: Metadata{},
	}
}

func expectedData() BoardWrapper {
	return BoardWrapper{
		Boards: []Board{
			{
				Name:    "D1",
				Vendor:  "A1",
				Core:    "C1",
				HasWifi: true,
			},
			{
				Name:    "A1",
				Vendor:  "v1",
				Core:    "C1",
				HasWifi: false,
			},
			{
				Name:    "Z1",
				Vendor:  "V1",
				Core:    "C1",
				HasWifi: false,
			},
		},
		Metadata: Metadata{
			Errors: Errors{
				HasErrors: false,
			},
			Totals: Totals{
				Vendors:     2,
				Boards:      3,
				WifiEnabled: 1,
			},
		},
	}
}
