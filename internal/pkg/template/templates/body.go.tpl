{{- /*gotype: github.com/gontainer/gontainer/internal/pkg/template.data*/ -}}

{{ $containerType := .Output.Meta.ContainerType }}

{{ if not .Stub }}
	{{template "params-comment" .Output.Params}}
	{{template "services-comment" .}}
{{ end }}

type {{$containerType}} struct {
	*{{ containerAlias }}.Container
}

func init() {
    interface_ := (*interface {
        // service container
        Get(serviceID string) (interface{}, error)
        GetInContext(ctx {{importAlias "context"}}.Context, serviceID string) (interface{}, error)
        CircularDeps() error
        OverrideService(serviceID string, s {{ containerAlias }}.Service)
        AddDecorator(tag string, decorator interface{}, deps ...{{ containerAlias }}.Dependency)
        IsTaggedBy(serviceID string, tag string) bool
        GetTaggedBy(tag string) ([]interface{}, error)
        GetTaggedByInContext(ctx {{importAlias "context"}}.Context, tag string) ([]interface{}, error)

        // param container
        GetParam(paramID string) (interface{}, error)
        OverrideParam(paramID string, d {{ containerAlias }}.Dependency)

        // misc
        HotSwap(func ({{ containerAlias }}.MutableContainer))
        Root() *{{ containerAlias }}.Container

        // getters
        {{ range $service := .Output.Services }}
            {{ if ne $service.Getter "" }}
                {{ $service.Getter }}() ({{ $service.Type }}, error)
                {{ $service.Getter }}InContext(ctx {{ importAlias "context" }}.Context) ({{ $service.Type }}, error)
                {{ if $service.MustGetter }}
                    Must{{ $service.Getter }}() {{ $service.Type }}
                    Must{{ $service.Getter }}InContext(ctx {{ importAlias "context" }}.Context) {{ $service.Type }}
                {{end}}
            {{ end }}
        {{end}}
    })(nil)

	var nilContainer *{{$containerType}}

	interfaceType := {{ importAlias "reflect" }}.TypeOf(interface_).Elem()
	implements := {{ importAlias "reflect" }}.TypeOf(nilContainer).Implements(interfaceType)

	if !implements {
		panic("generated container does not implement expected interface")
	}
}

{{template "container-getters" .}}

func {{ .Output.Meta.ContainerConstructor }}() ({{ if not .Stub}}rootGontainer{{end}} *{{$containerType}}) {
	{{- if .Stub }}
		panic("stub")
	{{- else }}
		{{template "container-constructor" .}}
	{{- end }}
}

{{ if not .Stub }}
// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *{{$containerType}}) _concatenateChunks(first func() (interface{}, error), chunks ...func() (interface{}, error)) (string, error) {
	r := ""
	for _, p := range append([]func() (interface{}, error){first}, chunks...) {
		chunk, err := p()
		if err != nil {
			return "", err
		}
		s, err := {{ exporterAlias }}.CastToString(chunk)
		if err != nil {
			return "", err
		}
		r += s
	}
	return r, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *{{$containerType}}) _paramTodo(params ...string) (interface{}, error) {
	if len(params) > 0 {
		return nil, {{ importAlias "errors" }}.New(params[0])
	}
	return nil, {{ importAlias "errors" }}.New("parameter todo")
}

{{ $envVarDoesntExist := "environment variable %+q does not exist" }}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *{{$containerType}}) _getEnv(key string, def ...string) (string, error) {
	val, ok := {{ importAlias "os" }}.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0], nil
		}
		return "", {{ importAlias "fmt" }}.Errorf({{ export $envVarDoesntExist }}, key)
	}
	return val, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *{{$containerType}}) _getEnvInt(key string, def ...int) (int, error) {
	val, ok := {{ importAlias "os" }}.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, {{ importAlias "fmt" }}.Errorf({{ export $envVarDoesntExist }}, key)
	}
	res, err := {{ importAlias "strconv" }}.Atoi(val)
	if err != nil {
		return 0, {{ importAlias "fmt" }}.Errorf("cannot cast env(%+q) to int: %w", key, err)
	}
	return res, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *{{$containerType}}) _callProvider(provider interface{}, args ...interface{}) (interface{}, error) {
	return {{ callerAlias }}.CallProvider(provider, args, true)
}
{{ end }}
