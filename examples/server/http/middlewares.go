package http

import (
	"fmt"
	"net/http"
	"time"
)

type logger interface {
	Println(v ...interface{})
}

func NewLoggerMiddleware(l logger, n http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		l.Println(fmt.Sprintf("%s %s", req.Method, req.RequestURI))
		n.ServeHTTP(res, req)
	})
}

func NewLogExecutionTimeMiddleware(id string, l logger, n http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		s := time.Now()
		n.ServeHTTP(res, req)
		l.Println(fmt.Sprintf("HTTP handler `%s`, execution time %s", id, time.Since(s)))
	})
}
