# Parameters

## Schema

The key-value objects with values of primitive types.

```yaml
parameters:
  nil:    ~                              # nil
  dbPort: 3306                           # int(3306)
  port:   '%envInt("APP_PORT", 9090)%'   # reads the env var "APP_PORT", if does not exist return 9090
  host:   'env("APP_HOST", "localhost")' # reads the env var "APP_HOST", if does not exist return "localhost"
  hostport: "%host%:%post%"              # gontainer.GetParam("host") + ":" + gontainer.GetParam("port")
```

The compiler finds chunks surrounded by `%`, and:

1. returns `%` for `%%`
2. refers to param `anotherParam` for `%anotherParam%`
3. invokes `env("APP_HOST")` for `%env("APP_HOST")%`

**Note**

For functions (e.g. `%env("APP_HOST")%`) the code between parentheses must be a valid GO code.
