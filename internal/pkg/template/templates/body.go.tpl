{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}

{{ $containerType := .Output.Meta.ContainerType }}

{{template "params-comment" .Output.Params}}
{{template "services-comment" .}}

type {{$containerType}} struct {
	*{{ containerAlias }}.SuperContainer
}

{{template "container-getters" .}}

func {{ .Output.Meta.ContainerConstructor }}() (rootGontainer interface{
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
	sc := {{ containerAlias }}.NewSuperContainer()
	rootGontainer = &{{$containerType}}{
		SuperContainer: sc,
	}

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
	concatenateChunks := func(first func() (interface{}, error), chunks ...func() (interface{}, error)) (string, error) {
		r := ""
		for _, p := range append([]func() (interface{}, error){first}, chunks...) {
			chunk, err := p()
			if err != nil {
				return "", err
			}
			s, err := {{ exporterAlias }}.ToString(chunk)
			if err != nil {
				return "", err
			}
			r += s
		}
		return r, nil
	}
	_ = concatenateChunks
	paramTodo := func(params ...string) (interface{}, error) {
		if len(params) > 0 {
			return nil, {{ importAlias "errors" }}.New(params[0])
		}
			return nil, {{ importAlias "errors" }}.New("parameter todo")
		}
	_ = paramTodo

	const envVarDoesntExist = "environment variable %+q does not exist"

	getEnv := func(key string, def ...string) (string, error) {
		val, ok := {{ importAlias "os" }}.LookupEnv(key)
		if !ok {
			if len(def) > 0 {
				return def[0], nil
			}
			return "", {{ importAlias "fmt" }}.Errorf(envVarDoesntExist, key)
		}
		return val, nil
	}
	_ = getEnv

	getEnvInt := func(key string, def ...int) (int, error) {
		val, ok := {{ importAlias "os" }}.LookupEnv(key)
		if !ok {
			if len(def) > 0 {
				return def[0], nil
			}
			return 0, {{ importAlias "fmt" }}.Errorf(envVarDoesntExist, key)
		}
		res, err := {{ importAlias "strconv" }}.Atoi(val)
		if err != nil {
			return 0, {{ importAlias "fmt" }}.Errorf("cannot cast env(%+q) to int: %w", key, err)
		}
		return res, nil
	}
	_ = getEnvInt

	getParam := func(n string) (interface{}, error) {
		return sc.GetParam(n)
	}
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
				s.ScopeDefault()
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
}
