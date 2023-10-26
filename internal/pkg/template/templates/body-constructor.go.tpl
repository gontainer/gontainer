{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}

{{define "container-constructor"}}
	{{ $containerType := .Output.Meta.ContainerType }}
	sc := &{{$containerType}}{
		Container: {{ containerAlias }}.New(),
	}
	rootGontainer = sc

	//
	//
	// #####################################
	// ############## Helpers ##############
	//
	//
	dependencyService := {{ containerAlias }}.NewDependencyService
	_ = dependencyService
	dependencyValue := {{ containerAlias }}.NewDependencyValue
	_ = dependencyValue
	dependencyTag := {{ containerAlias }}.NewDependencyTag
	_ = dependencyTag
	dependencyProvider := {{ containerAlias }}.NewDependencyProvider
	_ = dependencyProvider
	newService := {{ containerAlias }}.NewService
	_ = newService
	concatenateChunks := sc._concatenateChunks
	_ = concatenateChunks
	paramTodo := sc._paramTodo
	_ = paramTodo
	getEnv := sc._getEnv
	_ = getEnv
	getEnvInt := sc._getEnvInt
	_ = getEnvInt
	getParam := sc.GetParam
	_ = getParam

	{{ if .Output.Params }}
	//
	//
	// ####################################
	// ############## Params ##############
	//
	//
	{{ end }}
	{{ range $param := .Output.Params}}
	// {{ export $param.Name }}: {{ export $param.Raw }}
	sc.OverrideParam({{ export $param.Name }}, {{$param.Code}})
	{{ end}}

	{{ if .Output.Services }}
	//
	//
	// ######################################
	// ############## Services ##############
	//
	//
	{{ end }}
	{{range $service := .Output.Services}}
	// {{ export $service.Name }}
	{
		s := newService()
		{{ if eq $service.Todo true }}
			s.SetConstructor(func () (interface{}, error) { return nil, {{ importAlias "errors" }}.New("service todo") })
		{{ else }}
			{{ if ne $service.Constructor "" }}
				s.SetConstructor(
					{{ $service.Constructor }},
					{{ range $arg := $service.Args }}
						// {{ export $arg.Raw }}
						{{ $arg.Code }},
					{{end}}
				)
			{{ else if ne $service.Value "" }}
				s.SetConstructor(func () {{if ne $service.Type "" }} {{$service.Type}} {{else}} interface{} {{end}} { return {{ $service.Value }} })
			{{ else if ne $service.Type "" }}
				s.SetConstructor(func () (result {{ $service.Type }}) { return })
			{{ end }}

			{{ range $field := $service.Fields }}
				s.SetField({{ export $field.Name }}, {{ $field.Value.Code }} )
			{{end}}

			{{ range $call := $service.Calls }}
				{{ if eq $call.Immutable true }}
				s.AppendWither(
				{{else}}
				s.AppendCall(
				{{ end }}
					{{ export $call.Method }},
					{{ range $arg := $call.Args }}
						// {{ export $arg.Raw }}
						{{ $arg.Code }},
					{{end}}
				)
			{{end}}

			{{ range $tag := $service.Tags }}
				s.Tag({{ export $tag.Name }}, {{ export $tag.Priority }})
			{{ end }}

			{{ if $service.Scope.IsDefault }}
				s.SetScopeDefault()
			{{ else if $service.Scope.IsShared }}
				s.ScopeShared()
			{{ else if $service.Scope.IsContextual }}
				s.ScopeContextual()
			{{ else if $service.Scope.IsNonShared }}
				s.ScopeNonShared()
			{{ end }}
		{{ end }}
		sc.OverrideService({{ export $service.Name }}, s)
	}
	{{end}}

	{{ if .Output.Decorators }}
	//
	//
	// ########################################
	// ############## Decorators ##############
	//
	//
	{{ end }}
	{{ range $decorator := .Output.Decorators }}
		sc.AddDecorator(
			{{ export $decorator.Tag }},
			{{ $decorator.Decorator }},
			{{ range $arg := $decorator.Args }}
				{{ $arg.Code }},
			{{end}}
		)
	{{end}}

	return
{{- end }}
