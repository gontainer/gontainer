[![Build Status](https://github.com/gontainer/gontainer/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/gontainer/gontainer/actions?query=workflow%3ATests)
[![Coverage Status](https://coveralls.io/repos/github/gontainer/gontainer/badge.svg?branch=main)](https://coveralls.io/github/gontainer/gontainer?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gontainer/gontainer)](https://goreportcard.com/report/github.com/gontainer/gontainer)
[![Go Reference](https://pkg.go.dev/badge/github.com/gontainer/gontainer.svg)](https://pkg.go.dev/github.com/gontainer/gontainer)

# Gontainer

A Depenendency Injection container for GO inspired by [Symfony](https://symfony.com/doc/current/components/dependency_injection.html).

Using the bootstrapping technique, Gontainer uses itself to compile its dependencies.
1. [Configuration](gontainer)
2. [Usage](internal/cmd/runner_builder.go)

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

## TL;DR

**Describe dependencies**

Either YAML or GO

<details>
  <summary>YAML</summary>

File `gontainer/gontainer.yaml`:

```yaml
meta:
  pkg: "gontainer"
  constructor: "New"
  imports:
     mypkg: "github.com/user/repo/pkg"

parameters:
  appPort: '%envInt("APP_PORT", 9090)%' # get the port from the ENV variable if it exists, otherwise, use the default one

services:
  endpointHelloWorld:
    constructor: "mypkg.NewHelloWorld"

  serveMux:
    constructor: '"net/http".NewServeMux'                       # serveMux := http.NewServerMux()
    calls:                                                      #
      - [ "Handle", [ "/hello-world", "@endpointHelloWorld" ] ] # serveMux.Handle("/hello-world", gontainer.Get("endpointHelloWorld"))

  server:
    getter: "GetServer"           # func (*gontainer) GetServer() (*http.Server, error) { ... }
    must_getter: true             # func (*gontainer) MustGetServer() *http.Server { ... }
    type: '*"net/http".Server'    # 
    value: '&"net/http".Server{}' # server := &http.Server{}
    fields:                       #
      Addr: ":%appPort%"          # server.Addr = ":" + gontainer.GetParam("appPort")
      Handler: "@serveMux"        # server.Handler = gontainer.Get("serverMux")
```

**Compile it**

```bash
gontainer build -i gontainer/gontainer.yaml -o gontainer/container.go

# it can read multiple configuration files, e.g.
# gontainer build -i gontainer/gontainer.yaml -i gontainer/dev/\*.yaml -o gontainer/container.go
```
</details>

<details>
<summary>GO</summary>

File `gontainer/gontainer.go`:

```go
package gontainer

import (
   "net/http"
   "os"

   "github.com/gontainer/gontainer-helpers/container"
   "github.com/user/repo/pkg"
)

type gontainer struct {
   *container.SuperContainer
}

func (g *gontainer) MustGetServer() (r *http.Server) {
   err := g.CopyServiceTo("server", &r)
   if err != nil {
      panic(err)
   }
   return
}

func New() *gontainer {
   sc := &gontainer{
      SuperContainer: container.NewSuperContainer(),
   }

   sc.OverrideParam("serverAddr", container.NewDependencyProvider(func() string {
      if v, ok := os.LookupEnv("APP_PORT"); ok {
         return ":" + v
      }
      return ":9090"
   }))

   endpointHelloWorld := container.NewService()
   endpointHelloWorld.SetConstructor(pkg.NewHelloWorld)
   sc.OverrideService("endpointHelloWorld", endpointHelloWorld)

   serveMux := container.NewService()
   serveMux.SetConstructor(http.NewServeMux)
   serveMux.AppendCall("Handle", container.NewDependencyService("endpointHelloWorld"))
   sc.OverrideService("serveMux", serveMux)

   server := container.NewService()
   server.SetConstructor(func() *http.Server {
      return &http.Server{}
   })
   server.SetField("Addr", container.NewDependencyProvider(func() (interface{}, error) {
      return sc.GetParam("serverAddr")
   }))
   server.SetField("Handler", container.NewDependencyService("serveMux"))
   sc.OverrideService("server", server)

   return sc
}

```
</details>

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
