# Interface

Gontainer generates an object that implements the following interface
and contains additional getters defined in the provided configuration.

```go
import (
    "context"
	
    "github.com/gontainer/gontainer-helpers/v3/container"
)

type Container interface {
    // services
    Get(serviceID string) (interface{}, error)
    GetInContext(ctx context.Context, serviceID string) (interface{}, error)
    CircularDeps() error
    OverrideService(serviceID string, service container.Service)
    AddDecorator(tag string, decorator interface{}, deps ...container.Dependency)
    IsTaggedBy(serviceID string, tag string) bool
    GetTaggedBy(tag string) ([]interface{}, error)
    GetTaggedByInContext(ctx context.Context, tag string) ([]interface{}, error)
    // params
    GetParam(paramID string) (interface{}, error)
    OverrideParam(paramID string, d container.Dependency)
    // misc
    HotSwap(func (container.MutableContainer))
    Root() *container.Container
}
```
