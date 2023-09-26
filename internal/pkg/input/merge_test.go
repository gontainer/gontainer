package input_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	i1 := input.Input{
		Version: ptr.New(input.Version("0.1.0-alpha")),
		Meta: input.Meta{
			Pkg:                  nil,
			ContainerType:        ptr.New("MyGontainer"),
			ContainerConstructor: ptr.New("NewGontainer"),
		},
		Params: map[string]any{
			"app-name": "my-app",
			"host":     "localhost",
			"port":     3306,
		},
		Services: map[string]input.Service{
			"db": {
				Getter: ptr.New("GetDB"),
				Scope:  ptr.New(input.ScopeShared),
			},
			"server": {
				Constructor: ptr.New("NewServer"),
				Args:        []any{3306},
				Scope:       ptr.New(input.ScopePrivate),
			},
		},
	}
	i2 := input.Input{
		Meta: input.Meta{
			Pkg:           ptr.New("pkg"),
			ContainerType: nil,
		},
		Params: map[string]any{
			"port": 3307,
		},
		Services: map[string]input.Service{
			"logger": {
				Args: []any{"@stdout"},
			},
			"db": {
				Scope: ptr.New(input.ScopePrivate),
				Args:  []any{"localhost", 3306},
			},
			"server": {
				Scope: ptr.New(input.ScopeShared),
			},
		},
	}
	expected := input.Input{
		Version: ptr.New(input.Version("0.1.0-alpha")),
		Meta: input.Meta{
			Pkg:                  ptr.New("pkg"),
			ContainerType:        ptr.New("MyGontainer"),
			ContainerConstructor: ptr.New("NewGontainer"),
		},
		Params: map[string]any{
			"app-name": "my-app",
			"host":     "localhost",
			"port":     3307,
		},
		Services: map[string]input.Service{
			"db": {
				Getter: ptr.New("GetDB"),            // from i1
				Scope:  ptr.New(input.ScopePrivate), // from i2
				Args:   []any{"localhost", 3306},    // from i2
			},
			"logger": {
				Args: []any{"@stdout"},
			},
			"server": {
				Constructor: ptr.New("NewServer"),       // from i1
				Args:        []any{3306},                // from i1
				Scope:       ptr.New(input.ScopeShared), // from i2
			},
		},
	}

	merged := input.Merge(i1, i2)
	assert.Equal(t, expected.Version, merged.Version, "Version")
	assert.Equal(t, expected.Meta, merged.Meta, "Meta")
	assert.Equal(t, expected.Params, merged.Params, "Params")
	assert.Equal(t, expected.Services, merged.Services, "Services")
	assert.Equal(t, expected.Decorators, merged.Decorators, "Decorators")
}
