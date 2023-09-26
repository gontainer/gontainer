package input

import (
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/slices"
)

// newPtr creates a new pointer that points to a different address, but the same value.
func newPtr[T any](i *T) *T {
	if i == nil {
		return nil
	}
	o := *i
	return &o
}

func mergePtr[T any](a, b *T) *T {
	if b != nil {
		return newPtr(b)
	}
	return newPtr(a)
}

func mergeMap[T any](a, b map[string]T) map[string]T {
	if a == nil && b == nil {
		return nil
	}

	r := make(map[string]T)
	for _, m := range []map[string]T{a, b} {
		for k, v := range m {
			r[k] = v
		}
	}
	return r
}

func mergeMeta(m1, m2 Meta) Meta {
	return Meta{
		Pkg:                  mergePtr(m1.Pkg, m2.Pkg),
		ContainerType:        mergePtr(m1.ContainerType, m2.ContainerType),
		ContainerConstructor: mergePtr(m1.ContainerConstructor, m2.ContainerConstructor),
		DefaultMustGetter:    mergePtr(m1.DefaultMustGetter, m2.DefaultMustGetter),
		Imports:              mergeMap(m1.Imports, m2.Imports),
		Functions:            mergeMap(m1.Functions, m2.Functions),
	}
}

func mergeArgs(a, b []any) []any {
	if len(b) > 0 {
		return slices.Copy(b)
	}
	return slices.Copy(a)
}

func mergeService(s1, s2 Service) Service {
	return Service{
		Getter:      mergePtr(s1.Getter, s2.Getter),
		MustGetter:  mergePtr(s1.MustGetter, s2.MustGetter),
		Type:        mergePtr(s1.Type, s2.Type),
		Value:       mergePtr(s1.Value, s2.Value),
		Constructor: mergePtr(s1.Constructor, s2.Constructor),
		Args:        mergeArgs(s1.Args, s2.Args),
		Calls:       append(slices.Copy(s1.Calls), s2.Calls...),
		Fields:      mergeMap(s1.Fields, s2.Fields),
		Tags:        append(slices.Copy(s1.Tags), s2.Tags...),
		Scope:       mergePtr(s1.Scope, s2.Scope),
		Todo:        mergePtr(s1.Todo, s2.Todo),
	}
}

func mergeServices(a, b map[string]Service) map[string]Service {
	r := make(map[string]Service)

	maps.Iterate(a, func(k string, v1 Service) {
		r[k] = v1
	})
	maps.Iterate(b, func(k string, v2 Service) {
		if v1, ok := a[k]; ok {
			r[k] = mergeService(v1, v2)
			return
		}
		r[k] = v2
	})

	return r
}

func Merge(i1, i2 Input) Input {
	return Input{
		Version:    mergePtr(i1.Version, i2.Version),
		Meta:       mergeMeta(i1.Meta, i2.Meta),
		Params:     mergeMap(i1.Params, i2.Params),
		Services:   mergeServices(i1.Services, i2.Services),
		Decorators: append(slices.Copy(i1.Decorators), i2.Decorators...),
	}
}
