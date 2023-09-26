# TEST REPOSITORY; DO NOT USE IT

[![Build Status](https://github.com/gontainer/gontainer/workflows/Tests/badge.svg?branch=main)](https://github.com/gontainer/gontainer/actions?query=workflow%3ATests)
[![Coverage Status](https://coveralls.io/repos/github/gontainer/gontainer/badge.svg?branch=main)](https://coveralls.io/github/gontainer/gontainer?branch=main)

# Gontainer

A Depenendency Injection container for GO inspired by [Symfony](https://symfony.com/doc/current/components/dependency_injection.html).

## Docs

1. Documentation
   1. [Version](docs/VERSION.md)
   2. [Meta](docs/META.md)
   3. [Parameters](docs/PARAMETERS.md)
   4. [Services](docs/SERVICES.md)
   5. [Decorators](docs/DECORATORS.md)
2. Use cases
   1. [Composition root](docs/COMPOSITION_ROOT.md)
   2. [Contextual scope](docs/CONTEXTUAL_SCOPE.md)
3. [Examples](examples)
4. [Interface](docs/INTERFACE.md)

## Installation

**homebrew**

```bash
brew install gontainer/homebrew-tap/gontainer
```

**go install**

```bash
go install github.com/gontainer/gontainer@latest
```

**Manual compilation**

```bash
git clone git@github.com:gontainer/gontainer.git
cd gontainer
GONTAINER_BINARY=/usr/local/bin/gontainer make build
```

## TL;DR

**Describe dependencies in YAML**

File `gontainer/gontainer.yaml`:

```yaml
meta:
  pkg: "gontainer"
  constructor: "New"

parameters:
  appPort: '%envInt("APP_PORT", 9090)%' # get the port from the ENV variable if it exists, otherwise, use the default one

services:
  endpointHelloWorld: # sample HTTP endpoint
    constructor: "http.NewHelloWorld"

  serveMux:
    constructor: '"net/http".NewServeMux'
    calls:
      - [ "Handle", [ "/hello-world", "@endpointHelloWorld" ] ]

  server:
    getter: "GetServer"
    must_getter: true # define method MustGetServer
    value: '&"net/http".Server{}'
    type: '*"net/http".Server'
    fields:
      Addr: ":%appPort%"
      Handler: "@serveMux"
```

**Compile it**

```bash
gontainer build -i gontainer/gontainer.yaml -o gontainer/container.go

# it can read multiple configuration files, e.g.
# gontainer build -i gontainer/gontainer.yaml -i gontainer/dev/*.yaml -o gontainer/container.go
```

**Voilà!**

File `main.go`:

```go
package main

import (
	"github.com/user/repo/gontainer"
)

func main() {
	c := gontainer.New()
	s := c.MustGetServer()

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```
