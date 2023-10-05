{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}

{{ $containerType := .Output.Meta.ContainerType }}

{{ if not .Stub }}
	{{template "params-comment" .Output.Params}}
	{{template "services-comment" .}}
{{ end }}

type {{$containerType}} struct {
	*{{ containerAlias }}.SuperContainer
}

{{template "container-getters" .}}

func {{ .Output.Meta.ContainerConstructor }}() ({{ if not .Stub}}rootGontainer{{end}} interface{
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
		{{ $service.Getter }}() ({{ $service.Type }}, error)
		{{ if $service.MustGetter }}
			Must{{ $service.Getter }}() {{ $service.Type }}
		{{end}}
	{{ end }}
{{end}}
}) {
	{{- if .Stub }}
		panic("stub")
	{{- else }}
		{{template "container-constructor" .}}
	{{- end }}
}
