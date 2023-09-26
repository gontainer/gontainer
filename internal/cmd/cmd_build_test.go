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
	}

	runCmdScenarios(t, scenarios...)
}
