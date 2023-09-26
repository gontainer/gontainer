package input

type Input struct {
	Version    *Version           `yaml:"version"`
	Meta       Meta               `yaml:"meta"`
	Params     map[string]any     `yaml:"parameters"`
	Services   map[string]Service `yaml:"services"`
	Decorators []Decorator        `yaml:"decorators"`
}

type Version string

type Meta struct {
	Pkg                  *string           `yaml:"pkg"`                   // default "main"
	ContainerType        *string           `yaml:"container_type"`        // default "Gontainer"
	ContainerConstructor *string           `yaml:"container_constructor"` // default "NewGontainer"
	DefaultMustGetter    *bool             `yaml:"default_must_getter"`   // default false
	Imports              map[string]string `yaml:"imports"`               // {"alias": "my/long/path", ...}
	Functions            map[string]string `yaml:"functions"`             // {"env": "os.Getenv", ...}
}

type Scope uint

type Service struct {
	Getter      *string        `yaml:"getter"`      // e.g. GetDB
	MustGetter  *bool          `yaml:"must_getter"` // if true adds MustGet func, it requires non-empty Getter
	Type        *string        `yaml:"type"`        // *?my/import/path.Type
	Value       *string        `yaml:"value"`       // my/import/path.Variable
	Constructor *string        `yaml:"constructor"` // NewDB
	Args        []any          `yaml:"arguments"`   // ["%host%", "%port%", "@logger"]
	Calls       []Call         `yaml:"calls"`       // [["SetLogger", ["@logger"]], ...]
	Fields      map[string]any `yaml:"fields"`      // Field: "%value%"
	Tags        []Tag          `yaml:"tags"`        // ["service_decorator", ...]
	Scope       *Scope         `yaml:"scope"`       // scope for the service, if nil, default value will be applied
	Todo        *bool          `yaml:"todo"`        // if true skips validation and returns error whenever users asks container for a service
}

type Tag struct {
	Name     string
	Priority int
}

type Call struct {
	Method    string
	Args      []any
	Immutable bool
}

type Decorator struct {
	Tag       string `yaml:"tag"`       // http-client
	Decorator string `yaml:"decorator"` // myImport/pkgCopy.MakeTracedHttpClient
	Args      []any  `yaml:"arguments"` // ["@logger"]
}
