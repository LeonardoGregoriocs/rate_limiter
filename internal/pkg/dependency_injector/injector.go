package dependencyinjector

import (
	"time"

	"github.com/leonardogregoriocs/rate_limiter/config"
	"github.com/leonardogregoriocs/rate_limiter/internal/infra/database"
	"github.com/leonardogregoriocs/rate_limiter/internal/infra/web"
	"github.com/leonardogregoriocs/rate_limiter/internal/infra/web/handlers"
	middlewares "github.com/leonardogregoriocs/rate_limiter/internal/infra/web/middleware"
	logger "github.com/leonardogregoriocs/rate_limiter/internal/pkg/log"
	"github.com/leonardogregoriocs/rate_limiter/internal/pkg/ratelimiter"
	ratelimiter_strategies "github.com/leonardogregoriocs/rate_limiter/internal/pkg/ratelimiter/strategies"
	responsehandler "github.com/leonardogregoriocs/rate_limiter/internal/pkg/responsehandles"
)

type DependencyInjectorInterface interface {
	Inject() (*Dependencies, error)
}

type DependencyInjector struct {
	Config *config.Config
}

type Dependencies struct {
	Logger                logger.LoggerInterface
	ResponseHandler       responsehandler.WebResponseHandlerInterface
	HelloWebHandler       handlers.HelloWebHandlerInterface
	RateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface
	WebServer             web.WebServerInterface
	RedisDatabase         database.RedisDatabaseInterface
	RateLimiter           ratelimiter.RateLimiterInterface
	RedisLimiterStrategy  ratelimiter_strategies.LimiterStrategyInterface
}

func (d *DependencyInjector) Inject() (*Dependencies, error) {
	logger := logger.NewLogger(d.Config.LogLevel)
	responseHandler := responsehandler.NewWebResponseHandler()

	redisDB, err := database.NewConnection(*d.Config, logger.GetLogger())
	if err != nil {
		return nil, err
	}

	redisLimiterStrategy := ratelimiter_strategies.NewRedisLimiterStrategy(
		redisDB.Client,
		logger.GetLogger(),
		time.Now,
	)

	limiter := ratelimiter.NewRateLimiter(
		logger,
		redisLimiterStrategy,
		d.Config.RateLimiterIPMaxRequests,
		d.Config.RateLimiterTokenMaxRequests,
		d.Config.RateLimiterTimeWindowMilliseconds,
	)

	helloWebHandler := handlers.NewHelloWebHandler(responseHandler)
	rateLimiterMiddleware := middlewares.NewRateLimiterMiddleware(logger, responseHandler, limiter)

	webRouter := web.NewWebRouter(helloWebHandler, rateLimiterMiddleware)
	webServer := web.NewWebServer(
		d.Config.WebServerPort,
		logger.GetLogger(),
		webRouter.Build(),
		webRouter.BuildMiddlewares(),
	)

	return &Dependencies{
		Logger:                logger,
		ResponseHandler:       responseHandler,
		HelloWebHandler:       helloWebHandler,
		RateLimiterMiddleware: rateLimiterMiddleware,
		WebServer:             webServer,
		RedisDatabase:         redisDB,
		RateLimiter:           limiter,
		RedisLimiterStrategy:  redisLimiterStrategy,
	}, nil
}

func NewDependencyInjector(cfg *config.Config) *DependencyInjector {
	return &DependencyInjector{
		Config: cfg,
	}
}
