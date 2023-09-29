package cmd

import (
	"io"

	"github.com/gontainer/gontainer-helpers/container"
	"github.com/gontainer/gontainer/gontainer"
	"github.com/gontainer/gontainer/internal/cmd/runner"
)

func buildRunner(
	w io.Writer,
	version string,
	buildInfo string,
	paramsExistActive bool,
	servicesExistActive bool,
	inputPatterns []string,
	outputFile string,
	stub bool,
) *runner.Runner {
	c := gontainer.New()

	// @writer
	ws := container.NewService()
	ws.SetValue(w)
	c.OverrideService("writer", ws)

	// %version%
	c.OverrideParam("version", container.NewDependencyValue(version))

	// %buildInfo%
	c.OverrideParam("buildInfo", container.NewDependencyValue(buildInfo))

	// %inputPatterns%
	c.OverrideParam("inputPatterns", container.NewDependencyValue(inputPatterns))

	// %outputFile%
	c.OverrideParam("outputFile", container.NewDependencyValue(outputFile))

	// %stub%
	c.OverrideParam("stub", container.NewDependencyValue(stub))

	// method "Active" is over the value that decorator runner.DecorateStepVerboseSwitchable returns,
	// so it must be called manually, it cannot be registered in the YAML file
	//
	// calls:
	//   - [ "Active", [ "%paramsExistActive%" ] ]
	c.MustGetStepValidateParamsExist().Active(paramsExistActive)
	c.MustGetStepValidateServicesExist().Active(servicesExistActive)

	return c.MustGetRunner()
}
