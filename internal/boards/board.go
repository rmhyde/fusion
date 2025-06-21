package boards

import (
	"context"

	"github.com/rmhyde/fusion/internal/helpers"
	"github.com/rs/zerolog"
)

func (o Options) Combine(ctx context.Context) (BoardWrapper, error) {
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msgf("Getting json files from folder '%s'", o.Folder)
	files, err := o.getFilesRecursively(ctx)
	if err != nil {
		return BoardWrapper{}, err
	}

	logger.Debug().Msg("Reading in boards from files")
	wrapper := unmarshalFiles(ctx, files)

	logger.Debug().Msg("Sorting and gathering metrics on boards")

	wrapper = o.sortAndGatherMetrics(ctx, wrapper)

	return wrapper, nil
}

func unmarshalFiles(ctx context.Context, files []string) (wrapper BoardWrapper) {
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
