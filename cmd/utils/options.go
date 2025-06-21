package utils

import (
	"github.com/rmhyde/fusion/internal/boards"
	"github.com/spf13/cobra"
)

func NewCombineOptions(cmd *cobra.Command, args []string) (boards.Options, error) {
	options := boards.Options{
		Folder: ".",
	}

	if len(args) > 0 {
		options.Folder = args[0]
	}

	r, err := cmd.Flags().GetBool("recursive")
	if err != nil {
		return options, err
	}
	options.Recursive = r

	return options, nil
}
