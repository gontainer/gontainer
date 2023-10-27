package cmd

import (
	"io"

	"github.com/gontainer/gontainer-helpers/v2/container"
	"github.com/gontainer/gontainer/internal/cmd/runner"
	"github.com/gontainer/gontainer/internal/gontainer"
)

type runnerPayload struct {
	writer              io.Writer
	version             string
	buildInfo           string
	paramsExistActive   bool
	servicesExistActive bool
	inputPatterns       []string
	outputFile          string
	stub                bool
}

func buildRunner(p runnerPayload) *runner.Runner {
	c := gontainer.New()

	// @writer
	ws := container.NewService()
	ws.SetValue(p.writer)
	c.OverrideService("writer", ws)

	// %version%
	c.OverrideParam("version", container.NewDependencyValue(p.version))

	// %buildInfo%
	c.OverrideParam("buildInfo", container.NewDependencyValue(p.buildInfo))

	// %inputPatterns%
	c.OverrideParam("inputPatterns", container.NewDependencyValue(p.inputPatterns))

	// %outputFile%
	c.OverrideParam("outputFile", container.NewDependencyValue(p.outputFile))

	// %stub%
	c.OverrideParam("stub", container.NewDependencyValue(p.stub))

	// method "Active" is over the value that decorator runner.DecorateStepVerboseSwitchable returns,
	// so it must be called manually, it cannot be registered in the YAML file
	//
	// calls:
	//   - [ "Active", [ "%paramsExistActive%" ] ]
	c.MustGetStepValidateParamsExist().Active(p.paramsExistActive)
	c.MustGetStepValidateServicesExist().Active(p.servicesExistActive)

	return c.MustGetRunner()
}
