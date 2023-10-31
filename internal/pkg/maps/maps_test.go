// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
