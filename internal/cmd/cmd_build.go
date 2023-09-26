package cmd

import (
	"io"

	"github.com/fatih/color"
	"github.com/gontainer/gontainer-helpers/errors"
	"github.com/spf13/cobra"
)

func NewBuildCmd(version string, buildInfo string) *cobra.Command {
	var (
		inputPatterns         []string
		outputFile            string
		ignoreMissingParams   bool
		ignoreMissingServices bool
		quiet                 bool
		cmd                   *cobra.Command
	)

	cmd = &cobra.Command{
		Use:           "build",
		Short:         "Build a container",
		Long:          "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			out := io.Discard
			if !quiet {
				out = cmd.OutOrStdout()
			}

			err := buildRunner(
				out,
				version,
				buildInfo,
				!ignoreMissingParams,
				!ignoreMissingServices,
				inputPatterns,
				outputFile,
			).Run()

			if err == nil {
				return nil
			}

			errWriter := out

			_, _ = color.New(color.BgRed, color.FgHiWhite).Fprint(errWriter, "Errors:")
			_, _ = color.New().Fprintln(errWriter)
			for i, err := range errors.Collection(err) {
				_, _ = color.New().Fprint(errWriter, i+1, ". ")
				_, _ = color.New(color.BgRed, color.FgWhite).Fprint(errWriter, err)
				_, _ = color.New().Fprintln(errWriter)
			}

			return err
		},
	}

	cmd.Flags().StringArrayVarP(&inputPatterns, "input", "i", nil, "input name (required)")
	_ = cmd.MarkFlagRequired("input")

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "output name (required)")
	_ = cmd.MarkFlagRequired("output")

	cmd.Flags().BoolVarP(&ignoreMissingParams, "ignore-missing-params", "", false, "ignore missing parameters")
	cmd.Flags().BoolVarP(&ignoreMissingServices, "ignore-missing-services", "", false, "ignore missing services")
	cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "")

	return cmd
}
