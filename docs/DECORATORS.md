# Decorators

## Example

```yaml
services:
  logger:
    constructor: "log.Default"
    
  contactUsHandler:
    constructor: "NewContactUsHandler"

decorators:
  - tag: "http.handler"
    decorator: 'EndpointWithExecutionTime'
    arguments: ["@logger"]
```

A decorator must have at least 1 argument. The first one is always an instance of `container.DecoratorPayload`,
it holds the information about the used tag, the service name, and the created service we aim to decorate.
More arguments can be passed using `arguments`.

```go
// Middleware
func NewLogExecutionTimeMiddleware(id string, l logger, n http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		s := time.Now()
		n.ServeHTTP(res, req)
		l.Println(fmt.Sprintf("HTTP handler `%s`, execution time %s", id, time.Since(s)))
	})
}

// Decorator
func EndpointWithExecutionTime(payload container.DecoratorPayload, l *log.Logger) http.Handler {
	return NewLogExecutionTimeMiddleware(payload.ServiceID, l, payload.Service.(http.Handler))
}
```

```go
gontainer.Get("contactUsHandler")

// The above code is the equivalent of the following one:
//
// var handler interface{}
// handler = NewContactUsHandler()
// handler = EndpointWithExecutionTime(
//    container.DecoratorPayload{Tag: "http.handler", ServiceID: "contactUsHandler", Service: handler},
//    gontainer.Get("logger"),
// )
```
