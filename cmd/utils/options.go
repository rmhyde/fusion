package utils

import (
	"github.com/rmhyde/fusion/internal/boards"
	"github.com/rmhyde/fusion/internal/server"
	"github.com/spf13/cobra"
)

func NewCombineOptions(cmd *cobra.Command, args []string) (boards.Options, error) {
	options := boards.Options{
		Folder: ".",
		Ctx:    cmd.Context(),
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

func NewServeOptions(cmd *cobra.Command, args []string) (server.Options, error) {

	rootOptions, err := NewCombineOptions(cmd, args)
	options := server.Options{
		Ctx: cmd.Context(),
	}

	if err != nil {
		return options, err
	}

	options.Folder = rootOptions.Folder
	options.Recursive = rootOptions.Recursive

	p, err := cmd.Flags().GetInt("port")
	if err != nil {
		return options, err
	}
	options.Port = p

	h, err := cmd.Flags().GetString("host")
	if err != nil {
		return options, err
	}
	options.Host = h

	return options, nil
}
