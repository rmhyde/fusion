package boards

import (
	"context"
	"sort"
	"strings"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/rs/zerolog"
)

func (o Options) Combine(ctx context.Context) (BoardWrapper, error) {
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msgf("Getting json files from folder '%s'", o.Folder)
	files, err := o.getFiles(ctx)
	if err != nil {
		return BoardWrapper{}, err
	}

	logger.Debug().Msg("Reading in boards from files")
	wrapper := getBoards(ctx, files)

	logger.Debug().Msg("Sorting and gathering metrics on boards")

	wrapper = o.sortAndGatherMetrics(ctx, wrapper)

	return wrapper, nil
}

func getBoards(ctx context.Context, files []string) (wrapper BoardWrapper) {
	logger := zerolog.Ctx(ctx)
	for _, path := range files {
		b, err := helpers.ReadAsType[BoardWrapper](path)
		if err != nil {
			logger.Debug().Msgf("File: %s failed to read with error %s", path, err.Error())
			wrapper.Metadata.Errors.HasErrors = true
			wrapper.Metadata.Errors.FileReadErrors++
			wrapper.Metadata.Errors.Files = append(wrapper.Metadata.Errors.Files, path)
			continue
		}

		// This could be an issue reading all the boards into memory if there is a large number
		wrapper.Boards = append(wrapper.Boards, b.Boards...)
	}
	return
}

func (o Options) sortAndGatherMetrics(ctx context.Context, wrapper BoardWrapper) BoardWrapper {
	sort.Slice(wrapper.Boards, func(i, j int) bool {
		if wrapper.Boards[i].Vendor != wrapper.Boards[j].Vendor {
			return wrapper.Boards[i].Vendor < wrapper.Boards[j].Vendor
		}

		return wrapper.Boards[i].Name < wrapper.Boards[j].Name
	})

	vendors := make(map[string]int)
	for _, board := range wrapper.Boards {
		wrapper.Metadata.Totals.Boards++
		vendor := sanitizeVendor(board.Vendor)
		_, ok := vendors[vendor]
		if ok {
			vendors[vendor]++
		} else {
			vendors[vendor] = 1
		}
	}
	wrapper.Metadata.Totals.Vendors = len(vendors)
	return wrapper
}

func sanitizeVendor(vendor string) string {
	return strings.ToLower(vendor)
}
