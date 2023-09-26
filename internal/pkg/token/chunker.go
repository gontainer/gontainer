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
