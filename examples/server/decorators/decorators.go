package decorators

import (
	"log"
	"net/http"

	"github.com/gontainer/gontainer-helpers/container"
	pkgHttp "server/http"
)

func EndpointWithLogger(ctx container.DecoratorContext, l *log.Logger) http.Handler {
	return pkgHttp.NewLoggerMiddleware(l, ctx.Service.(http.Handler)) // todo do we need to have conversion here? maybe this context isn't the best idea
}

func EndpointWithExecutionTime(ctx container.DecoratorContext, l *log.Logger) http.Handler {
	return pkgHttp.NewLogExecutionTimeMiddleware(ctx.ServiceID, l, ctx.Service.(http.Handler))
}
