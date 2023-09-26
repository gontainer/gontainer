package maps_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	for i := 0; i < 100; i++ {
		require.Equal(
			t,
			[]string{"1", "2", "3"},
			maps.Keys(map[string]struct{}{"3": {}, "2": {}, "1": {}}),
		)
	}
}

func TestIterate(t *testing.T) {
	for i := 0; i < 100; i++ {
		var s []struct {
			Key string
			Val int
		}
		maps.Iterate(map[string]int{"2": 2, "3": 3, "1": 1}, func(k string, v int) {
			s = append(s, struct {
				Key string
				Val int
			}{Key: k, Val: v})
		})
		expected := []struct {
			Key string
			Val int
		}{
			{Key: "1", Val: 1},
			{Key: "2", Val: 2},
			{Key: "3", Val: 3},
		}
		require.Equal(t, expected, s)
	}
}
