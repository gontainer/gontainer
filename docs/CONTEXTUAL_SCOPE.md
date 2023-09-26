# Contextual scope

## Introduction

Gontainer lets the object be either `shared`, `non_shared`, or `contextual`.

* `shared` - caches objects, so the same instance will be used always
* `non_shared` - enforces Gontainer to create a new instance always
* `contextual` - shares the same instance only in the given context (the context is determined by a single invocation to the Gontainer, e.g. func `Get`)

**Default scope**

When the scope is not defined in the configuration, Gontainer determines the scope in the following way:

1. If the given object has at least one direct or indirect `contextual` dependency, the scope is `contextual`...
2. ...otherwise the scope is `shared`.

**Note**

Contextual dependencies may impact the scope of parent services.

### Example

In the following example the scope is not defined for `userStorage` neither for `userService`.
Gontainer will determine it automatically.
Since both of them have a `contextual` dependency, both of them will have the `contextual` scope.

```yaml
services:
  transaction:
    scope: "contextual"
  imageStorage:
    scope: "non_shared"
    constructor: "NewImageStorage"
    arguments: ["@transaction"]
  userStorage:
    # scope is not defined here, Gontainer will determine it automatically
    constructor: "NewUserStorage"
    arguments: ["@transaction", "@imageStorage", "@userStorage"]
  userService:
    # scope is not defined here, Gontainer will determine it automatically
    constructor: "NewUserService"
    arguments: ["@transaction"]
```

```
    gontainer.Get("userService") // single context
            |
            ↓
    userService (scope not defined)----
            |                         |
            |                         |---------------------------------------------|
            |                         ↓                                             ↓
            |                imageStorage (non_shared)                    userStorage (scope not defined)
            |                         |                                             |
            |                         |-----------|              |------------------|
            |                                     ↓              ↓
            |-------------------------------→ transaction (contextual)
                                                           ↓
                                                     db (shared)
```

```go
userService1 := gontainer.Get("userService")
userService2 := gontainer.Get("userService")

// false: each invocation creates a new instance
fmt.Println(userService1 == userService2)

// false: userService1 and userService1 have different transactions injected
fmt.Println(userService1.transaction == userService2.transaction)

// true: the transaction is shared between objects within the given context
fmt.Println(userService1.transaction == userService1.userStorage.transaction)
```

## Use case

To explain how to use the `contextual` scope, let's consider a simple HTTP endpoint that executes a few SQL queries in a transaction.

Requirements: write an endpoint that transfers funds between two accounts.

**Repositories**

```go
type AccountStorage struct {
	tx *sql.Tx
}

// Add adds funds to the given bank account. The `amount` can be negative.
func (a AccountStorage) Add(userID uint, amount int) error {
	_, err := a.tx.Exec("UPDATE accounts SET balance = balance + ? WHERE user_id = ?", amount, userID)
	return err
}

type AccountHistoryStorage struct {
	tx *sql.Tx
}

// AddTransaction adds a new entry to the history.
func (a AccountHistoryStorage) AddTransaction(userID uint, amount int, details string) error {
	_, err := a.tx.Exec("INSERT INTO transactions (user_id, amount, details), VALUES (?, ?, ?)", userID, amount, details)
	return err
}

type AccountFacade struct {
	storage        AccountStorage
	historyStorage AccountHistoryStorage
}

// Add adds a new entry to the history and updates the amount assigned to the given user.
func (a AccountFacade) Add(userID uint, amount int, details string) error {
	if err := a.storage.Add(userID, amount); err != nil {
		return fmt.Errorf("could not add funds: %w", err)
	}

	if err := a.historyStorage.AddTransaction(userID, amount, details); err != nil {
		return fmt.Errorf("could not add a new history entry: %w", err)
	}
	return nil
}
```

**Endpoint**

```go
func NewTransferHandler(accounts AccountFacade) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var (
			fromID uint
			toID   uint
			amount int
		)

		// extra logic

		if err := accounts.Add(fromID, -amount, fmt.Sprintf("transfer to %d", toID)); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := accounts.Add(toID, amount, fmt.Sprintf("transfer from %d", fromID)); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
	})
}
```

### Handling transactions

The above solution does not handle the transaction at all. To achieve it,
we would have to inject a new transaction at the beginning,
execute the business logic and roll it back or commit it depending on errors.
We may end up with something like:

```go
type TransactionFactory interface {
	New() *sql.Tx
}

func (a *AccountStorage) SetTransaction(tx *sql.Tx) {
	a.tx = tx
}

func (a *AccountHistoryStorage) SetTransaction(tx *sql.Tx) {
	a.tx = a.tx
}

func (a *AccountFacade) SetTransaction(tx *sql.Tx) {
	a.storage.SetTransaction(tx)
	a.historyStorage.SetTransaction(tx)
}

func NewTransferHandler(accounts AccountFacade, factory TransactionFactory) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var (
			fromID uint
			toID   uint
			amount int
			err    error
		)
		
		tx := factory.New()
		accounts.SetTransaction(tx)
		defer func() {
			if err == nil {
				tx.Commit()
				return
			}
			
			tx.Rollback()
		}()

		// extra logic
	})
}
```

The above solution requires that `AccountFacade` knows
that `AccountStorage` and `AccountHistoryStorage` need a transaction as well.

### Concurrency

Concurrent HTTP requests may share the same resources in GO. In the above case, they will use the same `AccountFacade`,
and it can lead to unexpected results when one request overrides a transaction started in another one.

We may inject a factory instead of a struct:

```go
type AccountFacadeFactory interface {
	CreateAccountFacade() *AccountFacade
}
```

Although in the end, we need to find a place where we will start a new transaction and wire all dependencies.
Gontainer solves this problem with the `contextual` scope.

### Solution

**Business logic**

```go
type HttpHandlerWithErrorFunc func(http.ResponseWriter, *http.Request) error

// NewTransferHandler does not need to know anything about the transaction
func NewTransferHandler(accounts AccountFacade) HttpHandlerWithErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) error {
		var (
			fromID uint
			toID   uint
			amount int
		)

		// extra logic

		if err := accounts.Add(fromID, -amount, fmt.Sprintf("transfer to %d", toID)); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return err
		}

		if err := accounts.Add(toID, amount, fmt.Sprintf("transfer from %d", fromID)); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return err
		}

		writer.WriteHeader(http.StatusOK)
		return nil
	}
}
```

**Decorator**

Notice that the following decorator can be reusable.

```go
// NewTransactionAwareEndpoint does not know anything about business logic
func NewTransactionAwareEndpoint(id string, h HttpHandlerWithErrorFunc, tx *sql.Tx) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := h(writer, request)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	})
}
```

**Decorate built-in ServeMux**

```go
type ServeMux struct {
	http.ServeMux
}

// Gontainer implements the following interface.
// Since it's not desired to rely on the specific type, let's define an interface.
type factory interface {
	Get(id string) (interface{}, error)
}

func (s ServeMux) HandleDynamic(pattern string, dynamicHandlerID string, f factory) {
	s.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Fetch the handler from the DI container and execute it.
		// Since our handler has the contextual scope, each time Gontainer will return a new instance.
		next, err := f.Get(dynamicHandlerID)
		if err != nil {
			panic(err)
		}
		next.(http.Handler).ServeHTTP(writer, request)
	}))
}
```

**Instruct Gointainer on how to create a transaction**

```go
func NewTx(db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}
```

**gontainer.yaml**

```yaml
meta:
  pkg: "main"
  container_type: "gontainer"
  container_constructor: "NewGontainer"

services:
  #################
  # DB components #
  #################

  db:
    constructor: "sql.Open"
    args: [ /* TODO */ ]

  transaction:
    scope: "contextual" # It must be contextual. Scopes of other objects are determined automatically.
    constructor: "NewTx"
    args: [ "@db" ]

  ##################
  # App components #
  ##################

  accountStorage:
    value: "AccountStorage{}"
    fields:
      tx: "@transaction"

  accountHistoryStorage:
    value: "AccountHistoryStorage{}"
    fields:
      tx: "@transaction"

  accountFacade:
    value: "AccountFacade{}"
    fields:
      storage: "@accountStorage"
      historyStorage: "@accountHistoryStorage"

  ##############
  # HTTP layer #
  ##############

  # endpoint /transfer
  endpointTransfer:
    constructor: "http.NewTransferHandler"
    arguments: [ "@facade" ]
    tags: [ "sql.transaction" ]

  serveMux:
    value: '&ServeMux{}'
    calls:
      # `$gontainer` is a special syntax, it allows access to the container itself
      - [ "HandleDynamic", [ "/transfer", "endpointTransfer", "$gontainer" ] ]

  server:
    getter: "GetServer"
    must_getter: true # create `MustGetServer` func
    value: '&"net/http".Server{}'
    type: '*"net/http".Server'
    fields:
      Addr: ":9090"
      Handler: "@serveMux"

decorators:
  # decorate all services tagged as "sql.transaction"
  - tag: "sql.transaction"
    decorator: "NewTransactionAwareEndpoint"
    # Decorators must have at least 2 arguments.
    # The first one is the ID of the service being decorated.
    # The second one is the service itself.
    # Next arguments must be defined as `arguments`.
    arguments: [ "@transaction" ]
```

**Voilà!**

```bash
gontainer build -i gontainer.yaml -o gontainer.go
```

```go
package main

func main() {
	g := NewGontainer()
	s := g.MustGetServer()

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```
