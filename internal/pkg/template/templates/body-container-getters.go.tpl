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
			    var s interface{}
				s, err = c.Get({{export $service.Name}})
				if err != nil {
                    return nil, {{ errorsAlias }}.PrefixedGroup(
                        {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
                        err,
                    )
                }
				err = {{ errorsAlias }}.PrefixedGroup(
                    {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
                    {{copierAlias}}.ConvertAndCopy(s, &result),
                )
                return
			{{- end }}
        }

		func (c *{{$containerType}}) {{ $service.Getter }}Context(ctx {{ importAlias "context" }}.Context) ({{ if not $stub }}result{{end}} {{ $service.Type }}, {{ if not $stub }}err{{ end }} error) {
			{{- if $stub }}
			    panic("stub")
			{{- else }}
                var s interface{}
                s, err = c.GetInContext(ctx, {{export $service.Name}})
                if err != nil {
                    return nil, {{ errorsAlias }}.PrefixedGroup(
                        {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
                        err,
                    )
                }
                err = {{ errorsAlias }}.PrefixedGroup(
                    {{ importAlias "fmt" }}.Sprintf("%s.%s(): ", {{export $containerType}}, {{export $service.Getter}}),
                    {{copierAlias}}.ConvertAndCopy(s, &result),
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

            func (c *{{$containerType}}) Must{{ $service.Getter }}Context(ctx {{ importAlias "context" }}.Context) {{ $service.Type }} {
                {{- if $stub }}
                    panic("stub")
                {{- else }}
                    r, err := c.{{ $service.Getter }}Context(ctx)
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
