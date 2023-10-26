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
			dependentTags     []string
			dependentServices []string
			dependentParams   []string
		)
		for _, arg := range s.AllArgs() {
			dependentServices = append(dependentServices, arg.DependsOnServices...)
			dependentTags = append(dependentTags, arg.DependsOnTags...)
			dependentParams = append(dependentParams, arg.DependsOnParams...)
		}
		g.ServiceDependsOnServices(s.Name, dependentServices)
		g.ServiceDependsOnTags(s.Name, dependentTags)
		g.ServiceDependsOnParams(s.Name, dependentParams)
	}
	for dID, d := range o.Decorators {
		g.AddDecorator(dID, d.Tag)
		var (
			dependentTags     []string
			dependentServices []string
			dependentParams   []string
		)
		for _, arg := range d.Args {
			dependentServices = append(dependentServices, arg.DependsOnServices...)
			dependentTags = append(dependentTags, arg.DependsOnTags...)
			dependentParams = append(dependentParams, arg.DependsOnParams...)
		}
		g.DecoratorDependsOnServices(dID, dependentServices)
		g.DecoratorDependsOnTags(dID, dependentTags)
		g.DecoratorDependsOnParams(dID, dependentParams)
	}

	for _, p := range o.Params {
		for _, p2 := range p.DependsOn {
			g.ParamDependsOnParam(p.Name, p2)
		}
	}
	return g
}
