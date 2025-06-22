package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rmhyde/fusion/cmd/serve"
	"github.com/rmhyde/fusion/cmd/utils"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:               "fusion",
	Short:             "CLI Tool for combining json files within the current or specified folder",
	PersistentPreRunE: rootCmdPersistentPreRunE,
	RunE:              runE,
}

func runE(cmd *cobra.Command, args []string) error {
	options, err := utils.NewCombineOptions(cmd, args)
	if err != nil {
		return err
	}

	wrapper, err := options.Combine()
	raw, err := json.Marshal(wrapper)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(string(raw))
	}
	return err
}

func Execute() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:          os.Stderr,
		PartsExclude: []string{zerolog.TimestampFieldName},
	}).With().Timestamp().Logger()

	cobra.CheckErr(RootCmd.ExecuteContext(logger.WithContext(context.Background())))
}

func init() {
	RootCmd.PersistentFlags().String("log-level", "info", "Log level")
	RootCmd.PersistentFlags().BoolP("recursive", "r", false, "Recursively get all json files within the parent folder and child folders")
	RootCmd.AddCommand(serve.ServeCmd)
}

func rootCmdPersistentPreRunE(cmd *cobra.Command, args []string) error {
	logger := zerolog.Ctx(cmd.Context())

	logLevelFlag, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return err
	}
	logLevel, err := zerolog.ParseLevel(logLevelFlag)
	if err != nil {
		logLevel = zerolog.InfoLevel
		logger.Warn().Err(err).Msg("couldn't parse log level, defaulting to INFO")
	}

	levelLogger := logger.Level(logLevel)
	levelLogger.With().Str("CommandName", cmd.CommandPath()).Logger()
	levelLogger.WithContext(cmd.Context())
	cmd.SetContext(levelLogger.WithContext(cmd.Context()))
	return nil
}
