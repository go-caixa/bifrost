package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/go-caixa/bifrost/common/logger"
	"github.com/go-caixa/bifrost/internal/config"
	"github.com/go-caixa/bifrost/internal/deliveries/healthz"
	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := logger.NewCtx(context.Background())

	env := flag.String("env", "local", "server environment")
	flag.Parse()

	logger.Infof(ctx, "initializing %s server", *env)
	conf := config.ReadConfig(ctx, *env)
	app := fiber.New()

	db := config.SetupDBConnection(ctx, *conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		logger.Infof(ctx, "Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	router := app.Group("/api/v1")
	healthz.Healthz(router.Group("/healthz"), conf, db)

	if err := app.Listen(conf.GetPort()); err != nil {
		logger.Fatalf(ctx, err, "failed starting the server at port: %s", conf.GetPort())
	}

	logger.Infof(ctx, "Running cleanup tasks...")
	logger.Infof(ctx, "closing db connection: %v", db.Close())
}
