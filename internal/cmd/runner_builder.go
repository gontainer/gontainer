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

	"github.com/gontainer/gontainer-helpers/v3/container"
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
