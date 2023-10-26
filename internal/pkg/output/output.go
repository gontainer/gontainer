package output

type Param struct {
	Name      string
	Code      string
	Raw       any
	DependsOn []string
}

type Arg struct {
	Code              string
	Raw               any
	DependsOnParams   []string
	DependsOnServices []string
	DependsOnTags     []string
}

type Call struct {
	Method    string
	Args      []Arg
	Immutable bool
}

type Field struct {
	Name  string
	Value Arg
}

type Tag struct {
	Name     string
	Priority int
}

type Scope uint

const (
	SetScopeDefault Scope = iota
	ScopeShared
	ScopeContextual
	ScopeNonShared
)

func (s Scope) IsDefault() bool {
	return s == SetScopeDefault
}

func (s Scope) IsShared() bool {
	return s == ScopeShared
}

func (s Scope) IsContextual() bool {
	return s == ScopeContextual
}

func (s Scope) IsNonShared() bool {
	return s == ScopeNonShared
}

type Service struct {
	Name        string
	Getter      string
	MustGetter  bool
	Type        string
	Value       string
	Constructor string
	Args        []Arg
	Calls       []Call
	Fields      []Field
	Tags        []Tag
	Scope       Scope
	Todo        bool
}

type Decorator struct {
	Tag       string
	Decorator string
	Args      []Arg
	Raw       string
}

type Meta struct {
	Pkg                  string
	ContainerType        string
	ContainerConstructor string
}

type Output struct {
	Meta       Meta
	Params     []Param
	Services   []Service
	Decorators []Decorator
}

// AllArgs returns arguments passed to constructor, calls and fields to fetch information about all dependencies.
// It does not include arguments passed to related decorators.
func (s Service) AllArgs() []Arg {
	var res []Arg
	res = append(res, s.Args...)
	for _, c := range s.Calls {
		res = append(res, c.Args...)
	}
	for _, f := range s.Fields {
		res = append(res, f.Value)
	}
	return res
}
