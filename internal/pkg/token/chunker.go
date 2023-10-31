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

package token

import (
	"fmt"
)

type Chunker struct {
}

func NewChunker() *Chunker {
	return &Chunker{}
}

// Chunks splits text into a slice. It finds substrings surrounded by percent marks ("%").
// An empty string returns a 1-length slice always.
//
//	chunker.Chunks("%firstname% %lastname%") // []string{"%firstname%", " ", "%lastname%"}
//	chunker.Chunks("")                       // []string{""}
func (c *Chunker) Chunks(s string) ([]string, error) {
	if s == "" {
		return []string{""}, nil
	}

	var r []string

	opened := false
	buff := ""

	for _, v := range s {
		ch := string(v)
		if ch == Delimiter {
			if opened {
				r = append(r, buff+Delimiter)
				buff = ""
				opened = false
				continue
			}

			if buff != "" {
				r = append(r, buff)
			}
			buff = Delimiter
			opened = true
			continue
		}

		buff += ch
	}

	if opened {
		return nil, fmt.Errorf("not closed token: %+q", buff)
	}

	if buff != "" {
		r = append(r, buff)
	}

	return r, nil
}
