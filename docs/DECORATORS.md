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

A decorator must have at least two arguments. The first one is always the name of the service that is being decorated.
The second one is always the service. More arguments can be passed using `arguments`.

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
func EndpointWithExecutionTime(id string, h http.Handler, l *log.Logger) http.Handler {
	return NewLogExecutionTimeMiddleware(id, l, h)
}
```

```go
gontainer.Get("contactUsHandler")

// The above code is the equivalent of the following one:
//
// handler := NewContactUsHandler()
// handler = EndpointWithExecutionTime("contactUsHandler", handler, gontainer.Get("logger"))
```
