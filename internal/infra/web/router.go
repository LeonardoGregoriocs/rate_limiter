package web

import (
	"github.com/leonardogregoriocs/rate_limiter/internal/infra/web/handlers"
	middlewares "github.com/leonardogregoriocs/rate_limiter/internal/infra/web/middleware"
)

type WebRouterInterface interface {
	Build() []RouteHandler
}

type WebRouter struct {
	HelloWebHandler       handlers.HelloWebHandlerInterface
	RateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface
}

func NewWebRouter(
	helloWebHandler handlers.HelloWebHandlerInterface,
	rateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface,
) *WebRouter {
	return &WebRouter{
		HelloWebHandler:       helloWebHandler,
		RateLimiterMiddleware: rateLimiterMiddleware,
	}
}

func (wr *WebRouter) Build() []RouteHandler {
	return []RouteHandler{
		{
			Path:        "/",
			Method:      "GET",
			HandlerFunc: wr.HelloWebHandler.SayHello,
		},
	}
}

func (wr *WebRouter) BuildMiddlewares() []Middleware {
	return []Middleware{
		{
			Name:    "RateLimiter",
			Handler: wr.RateLimiterMiddleware.Handle,
		},
	}
}
