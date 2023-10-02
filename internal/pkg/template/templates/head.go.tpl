{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}
{{ if .Stub }}
//go:build gontainerstub
// +build gontainerstub
{{ end }}
// Code generated by https://github.com/gontainer/gontainer; DO NOT EDIT.

package {{.Output.Meta.Pkg}}

// gontainer {{ .BuildInfo }}

import (
{{range $import := .ImportCollection.Imports -}}
    {{$import.Alias}} "{{$import.Path}}"
{{end}}
)