{{define "services-comment"}}

{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}
{{if .}}
// ············································································
// ··································SERVICES··································
// ············································································
{{end -}}

{{- $decorators := .Output.Decorators -}}

{{ range $service := .Output.Services -}}
    // #### {{$service.Name}}
    {{- if ne $service.Value "" -}}
        {{ "\n" }}// service := {{ $service.Value }}
    {{- else if ne $service.Type "" -}}
        {{ "\n" }}// var service {{ $service.Type }}
    {{- else -}}
        {{ "\n "}}// var service interface{}
    {{- end -}}

    {{- if ne $service.Constructor "" -}}
        {{ "\n" }}// service = {{$service.Constructor}}({{template "func-args" $service.Args}})
    {{- end -}}

    {{- range $call := $service.Calls -}}
        {{- if $call.Immutable -}}
            {{ "\n" }}// service = service.{{ $call.Method }}({{template "func-args" $call.Args}})
        {{- else -}}
            {{ "\n" }}// service.{{ $call.Method }}({{template "func-args" $call.Args}})
		{{- end -}}
    {{- end -}}

    {{- range $field := $service.Fields -}}
        {{ "\n" }}// {{ $service.Name }}.{{ $field.Name }} = {{template "export-raw" $field.Value.Raw}}
    {{- end -}}

	{{- range $decorator := $decorators -}}
	    {{- if isTagged $service.Name $decorator.Tag -}}
		    {{ "\n" }}// service = {{ $decorator.Decorator }}({{ export $service.Name }}, service{{ if $decorator.Args }}, {{template "func-args" $decorator.Args}}{{ end }})
		{{- end -}}
	{{- end -}}

    {{ "\n" }}// ············································································
{{end}}

{{end}}

{{define "export-raw"}}
		{{- if isString . -}}
		    eval({{- export . -}})
		{{- else -}}
		    {{- export . -}}
		{{- end -}}
{{end}}

{{define "func-args"}}
    {{- range $k, $arg := . -}}
        {{- if $k -}}
		    {{- ", " -}}
        {{- end -}}
		{{- template "export-raw" $arg.Raw -}}
    {{- end -}}
{{- end -}}
