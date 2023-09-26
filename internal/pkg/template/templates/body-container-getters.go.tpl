{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}
{{define "container-getters"}}

{{ $containerType :=.Output.Meta.ContainerType }}

{{ range $service := .Output.Services }}
    {{ if ne $service.Getter "" }}
        func (c *{{$containerType}}) {{ $service.Getter }}() (result {{ $service.Type }}, err error) {
			err = {{ errorsAlias }}.PrefixedGroup(
                {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
				c.CopyServiceTo({{export $service.Name}}, &result),
            )
            return
        }

        {{ if $service.MustGetter }}
            func (c *{{$containerType}}) Must{{ $service.Getter }}() {{ $service.Type }}{
                r, err := c.{{ $service.Getter }}()
                if err != nil {
                    panic(err.Error())
                }
                return r
            }
        {{ end }}
    {{ end }}
{{ end }}

{{end}}