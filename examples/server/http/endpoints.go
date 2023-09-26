package http

import (
	"fmt"
	"net/http"

	"github.com/gontainer/gontainer-helpers/container"
	"server/pkg"
)

func NewHomepage(desiredPath string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != desiredPath {
			http.NotFound(res, req)
			return
		}
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte("Welcome to the homepage"))
	})
}

func NewHelloWorld() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte("Hello World"))
	})
}

func NewContactUs() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte("contact@example.com"))
	})
}

type ServeMux struct {
	http.ServeMux
}

type factory interface {
	Get(id string) (interface{}, error)
}

func (s ServeMux) HandleDynamic(pattern string, dynamicHandlerID string, f factory) {
	s.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next, err := f.Get(dynamicHandlerID)
		if err != nil {
			panic(err)
		}
		next.(http.Handler).ServeHTTP(writer, request)
	}))
}

type HttpHandlerWithErrorFunc func(http.ResponseWriter, *http.Request) error

func NewTransactionAwareEndpoint(ctx container.DecoratorContext, t *pkg.Transaction) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx.Service.(HttpHandlerWithErrorFunc)(writer, request)
		_, _ = writer.Write([]byte("Func NewTransactionAwareEndpoint\n"))
		_, _ = writer.Write([]byte(fmt.Sprintf("Transaction.ID %s", t.ID)))
	})
}

func NewInTransactionHandler(u *pkg.Users) HttpHandlerWithErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) error {
		writer.Header().Set("Content-Type", "text/plain")
		_, _ = writer.Write([]byte("Func NewTransactionHandler\n"))
		_, _ = writer.Write([]byte(fmt.Sprintf("Users.Transaction.ID %s\n", u.Transaction.ID)))
		_, _ = writer.Write([]byte(fmt.Sprintf("Users.Accounts.Transaction.ID %s\n", u.Accounts.Transaction.ID)))
		_, _ = writer.Write([]byte("\n"))
		return nil
	}
}
