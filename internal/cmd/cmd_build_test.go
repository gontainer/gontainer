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

//go:build functional

package cmd_test

import (
	_ "embed"
	"testing"

	"github.com/gontainer/gontainer/internal/cmd"
	"github.com/spf13/cobra"
)

var (
	//go:embed testdata/circular-dep-all.txt
	buildCircularDepsAll string
	//go:embed testdata/invalid-input.txt
	buildInvalidInput string
	//go:embed testdata/circular-dep-params.txt
	buildCircularDepsParams string
	//go:embed testdata/circular-dep-services.txt
	buildCircularDepsServices string
	//go:embed testdata/matches-more-than-one.txt
	buildMatchesManyPatterns string
	//go:embed testdata/missing-params-and-services.txt
	buildMissingParamsAndServices string
	//go:embed testdata/missing-params-and-services-ignore-all.txt
	buildMissingParamsAndServicesIgnoreAll string
	//go:embed testdata/missing-params-and-services-ignore-params.txt
	buildMissingParamsAndServicesIgnoreParams string
	//go:embed testdata/func-does-not-exist.txt
	buildFuncDoesNotExist string
	//go:embed testdata/invalid-scope.txt
	buildInvalidScope string
	//go:embed testdata/invalid-tokens.txt
	buildInvalidTokens string
)

func TestNewBuildCmd(t *testing.T) {
	t.Parallel()

	newCmd := func() *cobra.Command {
		r := cmd.NewBuildCmd("dev-main", "test version")
		r.SilenceUsage = true
		r.SilenceErrors = true
		return r
	}

	scenarios := []cmdScenario{
		{
			cmd:   newCmd(),
			args:  "-i testdata/circular-dep-params.yaml -o /dev/null",
			out:   buildCircularDepsParams,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/circular-dep-services.yaml -o /dev/null",
			out:   buildCircularDepsServices,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/circular-dep-*.yaml -o /dev/null",
			out:   buildCircularDepsAll,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i [] -o /dev/null",
			out:   buildInvalidInput,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/circular-dep-services.yaml -i testdata/circular-*-services.yaml -o /dev/null",
			out:   buildMatchesManyPatterns,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/missing-params-and-services.yaml -o /dev/null",
			out:   buildMissingParamsAndServices,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/missing-params-and-services.yaml -o /dev/null --ignore-missing-params --ignore-missing-services",
			out:   buildMissingParamsAndServicesIgnoreAll,
			error: false,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/missing-params-and-services.yaml -o /dev/null --ignore-missing-params",
			out:   buildMissingParamsAndServicesIgnoreParams,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/func-does-not-exist.yaml -o /dev/null",
			out:   buildFuncDoesNotExist,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/invalid-scope.yaml -o /dev/null",
			out:   buildInvalidScope,
			error: true,
		},
		{
			cmd:   newCmd(),
			args:  "-i testdata/invalid-tokens.yaml -o /dev/null",
			out:   buildInvalidTokens,
			error: true,
		},
	}

	runCmdScenarios(t, scenarios...)
}
