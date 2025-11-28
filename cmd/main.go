package main

import (
	"context"
	"errors"
	"github.com/Mohamadreza-shad/auth/config"
	"github.com/Mohamadreza-shad/auth/internal/app/adapters/handlers/health"
	"github.com/Mohamadreza-shad/auth/internal/app/adapters/handlers/http"
	"github.com/Mohamadreza-shad/auth/pkg/i18n"
	"github.com/Mohamadreza-shad/auth/pkg/logging"
	"github.com/Mohamadreza-shad/auth/pkg/logging/keyval"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fx.New(
		fx.Provide(
			config.GetConfig,
			logging.NewLogger,
			i18n.NewI18n,
			http.NewGoHexagonalHttpHandler,
			http.NewRouter,
			health.NewHealthRouter,
		),
		fx.WithLogger(func(logger logging.Logger) fxevent.Logger {
			return logging.NewFxLogger(logger)
		}),
		fx.Invoke(
			startHTTPServers,
			handleShutdown,
		),
	).Run()
}

func startHTTPServers(lc fx.Lifecycle, cfg *config.Config, logger logging.Logger, router *http.Router, healthRouter *health.HealthRouter) {
	httpServer := &nethttp.Server{
		Addr:    cfg.Server.HttpAddress,
		Handler: router.Handler,
	}
	healthCheckServer := &nethttp.Server{
		Addr:    cfg.Server.HealthAddress,
		Handler: healthRouter.Handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("starting HTTP server: " + cfg.Server.HttpAddress)
				if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, nethttp.ErrServerClosed) {
					logger.Error("HTTP server stopped unexpectedly", keyval.Error(err))
				}
			}()
			go func() {
				logger.Info("starting health check server on: " + cfg.Server.HealthAddress)
				if err := healthCheckServer.ListenAndServe(); err != nil && !errors.Is(err, nethttp.ErrServerClosed) {
					logger.Error("Health check server stopped unexpectedly", keyval.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server")
			err := httpServer.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			logger.Info("Shutting down Health Check server")
			err = healthCheckServer.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			return nil
		},
	})
}

func handleShutdown(lc fx.Lifecycle, logger logging.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				sig := make(chan os.Signal, 1)
				signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
				<-sig
				logger.Info("Shutdown signal received")
				os.Exit(0)
			}()
			return nil
		},
	})
}
