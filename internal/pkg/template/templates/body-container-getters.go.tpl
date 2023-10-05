{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}
{{define "container-getters"}}

{{ $containerType :=.Output.Meta.ContainerType }}
{{ $stub := .Stub }}

{{ range $service := .Output.Services }}
    {{ if ne $service.Getter "" }}
        func (c *{{$containerType}}) {{ $service.Getter }}() ({{ if not $stub }}result{{end}} {{ $service.Type }}, {{ if not $stub }}err{{ end }} error) {
			{{- if $stub }}
			    panic("stub")
			{{- else }}
                err = {{ errorsAlias }}.PrefixedGroup(
                    {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
                    c.CopyServiceTo({{export $service.Name}}, &result),
                )
                return
			{{- end }}
        }

        {{ if $service.MustGetter }}
            func (c *{{$containerType}}) Must{{ $service.Getter }}() {{ $service.Type }}{
				{{- if $stub }}
				    panic("stub")
				{{- else }}
                    r, err := c.{{ $service.Getter }}()
                    if err != nil {
                        panic(err.Error())
                    }
                    return r
                {{- end }}
            }
        {{ end }}
    {{ end }}
{{ end }}

{{end}}
