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
