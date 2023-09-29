package templates

import (
	"embed"
)

var (
	//go:embed body*.go.tpl
	Body embed.FS
	//go:embed stub-body.go.tpl
	BodyStub embed.FS
	//go:embed head*.go.tpl
	Head embed.FS
)
