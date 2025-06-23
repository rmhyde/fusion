package boards

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
)

func (o Options) getFilesRecursively() (jsonFiles []string, err error) {
	logger := zerolog.Ctx(o.Ctx)
	logger.Debug().Msgf("Verifying folder '%s'", o.Folder)

	// Verify folder
	_, err = os.Stat(o.Folder)
	if err != nil {
		return
	}

	err = filepath.WalkDir(o.Folder, func(path string, d fs.DirEntry, err error) error {
		if !o.Recursive && d.IsDir() && path != o.Folder {
			logger.Debug().Msg("Recursion is not enabled and this isn't the parent directory, skipping directory")
			return filepath.SkipDir
		}

		logger.Debug().Msgf("Path: %s, d: %s, err: %v", path, d, err)
		// Currently we will just check for .json files but could always improve with a wildcard feature
		if !strings.HasSuffix(strings.ToLower(path), "json") {
			logger.Debug().Msgf("Skipping %s as extension is not valid", path)
			return nil
		}
		jsonFiles = append(jsonFiles, path)
		return nil
	})

	return
}
