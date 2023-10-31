// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"io"

	"github.com/fatih/color"
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/spf13/cobra"
)

func NewBuildCmd(version string, buildInfo string) *cobra.Command {
	var (
		inputPatterns         []string
		outputFile            string
		ignoreMissingParams   bool
		ignoreMissingServices bool
		quiet                 bool
		stub                  bool
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

			err := buildRunner(runnerPayload{
				writer:              out,
				version:             version,
				buildInfo:           buildInfo,
				paramsExistActive:   !ignoreMissingParams,
				servicesExistActive: !ignoreMissingServices,
				inputPatterns:       inputPatterns,
				outputFile:          outputFile,
				stub:                stub,
			}).Run()

			if err == nil {
				return nil
			}

			errWriter := out

			_, _ = color.New(color.BgRed, color.FgHiWhite).Fprint(errWriter, "Errors:")
			_, _ = color.New().Fprintln(errWriter)
			for i, err := range grouperror.Collection(err) {
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
	cmd.Flags().BoolVarP(&stub, "stub", "", false, "generate a stub")

	return cmd
}
