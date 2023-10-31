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

package output

import (
	"github.com/gontainer/gontainer-helpers/v2/container/graph"
)

type dependencyGraph interface {
	CircularDeps() [][]graph.Dependency       //nolint:staticcheck
	Deps(serviceID string) []graph.Dependency //nolint:staticcheck
}

func (o Output) BuildDependencyGraph() dependencyGraph {
	g := graph.New() //nolint:staticcheck
	for _, s := range o.Services {
		tags := make([]string, len(s.Tags))
		for i, t := range s.Tags {
			tags[i] = t.Name
		}
		g.AddService(s.Name, tags)

		var (
			dependantTags     []string
			dependantServices []string
			dependantParams   []string
		)
		for _, arg := range s.AllArgs() {
			dependantServices = append(dependantServices, arg.DependsOnServices...)
			dependantTags = append(dependantTags, arg.DependsOnTags...)
			dependantParams = append(dependantParams, arg.DependsOnParams...)
		}
		g.ServiceDependsOnServices(s.Name, dependantServices)
		g.ServiceDependsOnTags(s.Name, dependantTags)
		g.ServiceDependsOnParams(s.Name, dependantParams)
	}
	for dID, d := range o.Decorators {
		g.AddDecorator(dID, d.Tag)
		var (
			dependantTags     []string
			dependantServices []string
			dependantParams   []string
		)
		for _, arg := range d.Args {
			dependantServices = append(dependantServices, arg.DependsOnServices...)
			dependantTags = append(dependantTags, arg.DependsOnTags...)
			dependantParams = append(dependantParams, arg.DependsOnParams...)
		}
		g.DecoratorDependsOnServices(dID, dependantServices)
		g.DecoratorDependsOnTags(dID, dependantTags)
		g.DecoratorDependsOnParams(dID, dependantParams)
	}

	for _, p := range o.Params {
		for _, p2 := range p.DependsOn {
			g.ParamDependsOnParam(p.Name, p2)
		}
	}
	return g
}
