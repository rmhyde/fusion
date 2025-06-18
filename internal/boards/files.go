package boards

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
)

func (o Options) getFiles(ctx context.Context) (jsonFiles []string, err error) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msgf("Verifying folder '%s'", o.Folder)

	// Verify folder
	_, err = os.Stat(o.Folder)
	if err != nil {
		return
	}

	allFiles, err := os.ReadDir(o.Folder)
	if err != nil {
		logger.Err(err).Msgf("Error reading from %s", o.Folder)
		return
	}

	for _, file := range allFiles {
		if file.IsDir() {
			continue
		}

		logger.Debug().Msg(file.Name())
		// Currently we will just check for .json files but could always improve with a wildcard feature
		if !strings.HasSuffix(strings.ToLower(file.Name()), "json") {
			logger.Debug().Msgf("Skipping %s as extension is not valid", file.Name())
			continue
		}

		path := filepath.Join(o.Folder, file.Name())
		jsonFiles = append(jsonFiles, path)
	}
	return
}
