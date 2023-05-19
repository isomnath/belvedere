package web

import (
	"net/http"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/handlers"

	ddTracerMux "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Route struct {
	Name            string
	Path            string
	Middlewares     []Middleware
	HandlerFunction http.HandlerFunc
	Method          string
}

type Middleware struct {
	Function func(http.HandlerFunc) http.HandlerFunc
}

// Initialize - Returns a fully qualified mux router
func Initialize(routes []Route) *ddTracerMux.Router {
	router := ddTracerMux.NewRouter()

	router.HandleFunc(config.GetAppHealthCheckAPIPath(), baseResponseHeaderMiddleware(handlers.Ping)).Methods(http.MethodGet)
	router.NotFoundHandler = baseResponseHeaderMiddleware(handlers.RouteNotFoundHandler)

	for _, route := range routes {
		composedHandlerFunc := attachMiddlewares(route.Middlewares, route.HandlerFunction)
		router.HandleFunc(route.Path, baseResponseHeaderMiddleware(composedHandlerFunc)).Methods(route.Method)
	}

	return router
}

func attachMiddlewares(middlewares []Middleware, handlerFunc http.HandlerFunc) http.HandlerFunc {
	if middlewares == nil || len(middlewares) < 1 {
		return handlerFunc
	}
	composedHandlerFunc := handlerFunc
	for _, middleware := range middlewares {
		composedHandlerFunc = attachMiddleware(middleware, composedHandlerFunc)
	}
	return composedHandlerFunc
}

func attachMiddleware(middleware Middleware, handlerFunc http.HandlerFunc) http.HandlerFunc {
	composedHandlerFunc := middleware.Function(handlerFunc)
	return composedHandlerFunc
}

func baseResponseHeaderMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		span, ctx := tracer.StartSpanFromContext(request.Context(), "http.request", tracer.ResourceName(request.URL.Path))
		request.WithContext(ctx)
		user := request.Header.Get("username")
		tracer.SetUser(span, user, tracer.WithUserEmail(user))
		writer.Header().Set("Content-Type", "application/json")
		next(writer, request)
		span.Finish()
	}
}
