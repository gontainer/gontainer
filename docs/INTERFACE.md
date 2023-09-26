# Interface

Gontainer generates an object that implements the following interface
and contains additional getters defined in the provided configuration.

```go
import (
    "github.com/gontainer/gontainer-helpers/container"
)

type Container interface {
    // services
    Get(serviceID string) (result interface{}, err error)
    CircularDeps() error
    OverrideService(serviceID string, service container.Service)
    AddDecorator(tag string, decorator interface{}, deps ...container.Dependency)
    IsTaggedBy(serviceID string, tag string) bool
    GetTaggedBy(tag string) ([]interface{}, error)
    CopyServiceTo(serviceID string, dst interface{}) error

    // params
    GetParam(paramID string) (interface{}, error)
    OverrideParam(paramID string, d container.Dependency)
}
```
