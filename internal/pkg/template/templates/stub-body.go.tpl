{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}

{{ $containerType := .Output.Meta.ContainerType }}

type {{$containerType}} struct {
	*{{ containerAlias }}.SuperContainer
}

{{ range $service := .Output.Services }}
    {{ if ne $service.Getter "" }}
        func (c *{{$containerType}}) {{ $service.Getter }}() ({{ $service.Type }}, error) {
			panic("stub")
        }

        {{ if $service.MustGetter }}
            func (c *{{$containerType}}) Must{{ $service.Getter }}() {{ $service.Type }}{
				panic("stub")
            }
        {{ end }}
    {{ end }}
{{ end }}

func {{ .Output.Meta.ContainerConstructor }}() (interface{
	// service container
	Get(serviceID string) (interface{}, error)
	CircularDeps() error
	OverrideService(serviceID string, s {{ containerAlias }}.Service)
	AddDecorator(tag string, decorator interface{}, deps ...{{ containerAlias }}.Dependency)
	IsTaggedBy(serviceID string, tag string) bool
	GetTaggedBy(tag string) ([]interface{}, error)
	CopyServiceTo(serviceID string, dst interface{}) error

	// param container
	GetParam(paramID string) (interface{}, error)
	OverrideParam(paramID string, d {{ containerAlias }}.Dependency)

	// getters
{{ range $service := .Output.Services }}
	{{ if ne $service.Getter "" }}
		{{ $service.Getter }}() (result {{ $service.Type }}, err error)
		{{ if $service.MustGetter }}
			Must{{ $service.Getter }}() {{ $service.Type }}
		{{end}}
	{{ end }}
{{end}}
}) {
	panic("stub")
}
