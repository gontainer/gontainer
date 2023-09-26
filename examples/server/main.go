package main

import (
	"server/gontainer"
)

func main() {
	c := gontainer.New()
	s := c.MustGetServer()

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
