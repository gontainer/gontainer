package compiler

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
)

const (
	// meta
	defaultMetaPkg                  = "main"
	defaultMetaContainerType        = "Gontainer"
	defaultMetaContainerConstructor = "NewGontainer"
	defaultMetaMustGetter           = false

	// services
	defaultServiceGetter = ""
	defaultServiceTodo   = input.DefaultServiceTodo
)
