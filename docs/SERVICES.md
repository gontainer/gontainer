# Services

1. [Schema](#schema)
2. [Attributes](#attributes)
3. [Arguments and fields](#arguments-and-fields)
4. [Creating a new service](#creating-a-new-service)

## Schema

```yaml
services:
  serviceName:
    getter:      "GetService"
    must_getter: true
    type:        '*"net/http".Server'
    value:       '&"net/http".Server{}'
    constructor: '"my/import".NewServer'
    arguments:   ["%appPort%"]
    calls:
      - ["SetTimeout", ["%timeout%"]]
      - ["WithTimeout", ["%timeout%"], true]
    fields:
      Timeout: "%timeout%"
    tags:  ["my-name", {"name": "another-tag", "priority": 50}]
    scope: "shared" # enum("shared", "contextual", "non_shared")
    todo: false
```

## Attributes

**getter**

The func name for the getter. Default empty.

```yaml
services:
  serveMux:
    value: '&"net/http".ServeMux{}'
    getter: "GetServeMux"
```

```go
// func (g *Gontainer) GetServeMux() (interface{}, error)
serveMux, err := gontainer.GetServeMux()
```

**must_getter**

It requires a valid getter.

```yaml
services:
  serveMux:
    value:       '&"net/http".ServeMux{}'
    getter:      "GetServeMux"
    must_getter: true
```

```go
serveMux := gontainer.MustGetServeMux()
```

**type**

```yaml
services:
  serveMux:
    value:  '&"net/http".ServeMux{}'
    getter: "GetServeMux"
    type:   '*"net/http".Server'
```

```go
// Gontainer uses the type for generating getters. Whenever the type is empty
// interface{} will be used.
//
// func (g *Gontainer) GetServeMux() (*http.Server, error)
serveMux, err := gontainer.GetServeMux()
```

**value**

It accepts the same syntax as `!value ` (but the `!value ` prefix must not be given) see [arguments and fields](#arguments-and-fields).

```yaml
services:
  serveMux:
    value:  '&"net/http".ServeMux{}'
```

The service is initiated by the provided value.

**constructor**

```yaml
services:
  serveMux:
    constructor:  '"my/import".NewServer'
```

Constructor must refer to a valid func. The given func must return a single or two value.
If the func returns two values, the second value should be of type error.
If the second value is not nil, the error will be reported.

**arguments**

```yaml
services:
  logger:
    constructor: "NewLogger"
  db:                                         # It is the equivalent of the following code:
    constructor:  'NewDB'                     #
    arguments: ["localhost", 3306, "@logger"] # NewDB("localhost", 3306, gontainer.Get("logger"))
```

**calls**

A single call is an array of 3 elements:

1. Method name
2. Arguments
3. Is it a wither or a getter (`false`: getter, `true`: wither, default `false`)

```yaml
services:
  httpClient:
    value: '&"my/pkg/http".Client{}'         # service := &http.Client{}
    calls:                                   #
      - ["SetTimeout", ["%timeout%"]]        # service.SetTimeout(gontainer.GetParam("timeout"))
      - ["WithTimeout", ["%timeout%"], true] # service = service.WithTimeout(gontainer.GetParam("timeout"))
```

**fields**

Allows for the field injection (including unexported fields).

```yaml
services:
  httpClient:
    value: '&"my/pkg/http".Client{}' # service := &http.Client{}
    fields:                          #
      timeout: "%timeout%"           # service.timeout = gontainer.GetParam("timeout")
```

**tags**

Allows for tagging services. Tags can have priorities (used for `!tagged` syntax).

```yaml
services:
  serveMux:
    value:  '&"net/http".ServeMux{}'
    tags: ["my-tag", {"name": "another-tag", "priority": 50}]
```

**scope**

Defines whether the service can be cached and reused (`shared`) or no (`non_shared`).
The third value (`contextual`) is documented [here](CONTEXTUAL_SCOPE.md).

**todo**

Gontainer returns an error for any service with the flag `todo: true` always.

```yaml
services:
  logger:
    todo: true

  db:
    constructor: "NewDB"
    arguments: ["@logger"]
```

```go
db, err := gontainer.Get("db")
fmt.Println(err) // error: service todo

gontainer.Override("logger", ...) // define a proper logger here
db, err = gontainer.Get("db")
fmt.Println(err) // nil
```

## Arguments and fields

Arguments and fields supports a special syntax. That allows for injecting services and parameters.

1. `!value MyValue` -
    this syntax allows for injecting existing variables:
    * `!value MyVar`
    * `!value *MyVar`
    * `!value "my/pkg".Value`
    * `!value pkg.Value`
    * `!value "pkg".MyVar.Field` - `pkg` must be surrounded by `"` to inform the compiler which part points to the package
    * `!value &"pkg".MyVar.Field`
    * `!value ".".MyVar.Field` - `"."` informs the compiler that the variable `MyVar` exists in the current package
    * `!value "my/pkg".MyStruct{}`
    * `!value &"my/pkg".MyStruct{}`
    * `!value &MyStruct{}`
2. `@serviceName` -
   this syntax allows for referring to another service with the given name.
3. `!tagged tagName` - slice of all services tagged with `tagName`, sorted by priority descending then lexically by name whenever priorities are equal
4. `$gontainer` - it is a special variable, it allows for injecting the generated Gontainer
5. `%host%:%envInt("PORT")%` - whenever the argument does not match any of the above patterns and is a string,
it will be processed in the same way as a [parameter](PARAMETERS.md).

## Creating a new service

Gontainer gives 3 options to create a new service: using a `constructor`, using a `type`, or using a specific `value`.
One service cannot have a `constructor` and a `value`, because they mutually exclude each other.
Since Gontainer uses `type` to generate getters, `type` does not exclude `constructor` and `value`.

**constructor**

```yaml
services:
  db:
    constructor: "NewDB"                                   # service := NewDB(
    arguments: ['%env("DB_HOST")%', '%envInt("DB_PORT")%'] #     getEnv("DB_HOST"),
                                                           #     getEnv("DB_PORT"),
                                                           # )
```

**value**

```yaml
services:
  logger:
    value: "pkg.Logger{}" # service := pkg.Logger{}
```

**type**

```yaml
services:
  httpClient:
    type: "pkg.HttpClient"
    fields:                # var service pkg.HttpClient
      timeout: 3600        # service.timeout = 3600
```
