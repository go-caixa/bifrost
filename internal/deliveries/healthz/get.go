package healthz

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-caixa/bifrost/internal/config"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/hellofresh/health-go/v5"
)

func Healthz(r fiber.Router, cfg *config.Configuration, db *sql.DB) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    cfg.AppName,
		Version: cfg.GetAppVersion(),
	}), health.WithChecks(health.Config{
		Name:      "bifrost-db",
		Timeout:   30 * time.Second,
		SkipOnErr: false,
		Check: func(ctx context.Context) error {
			return db.Ping()
		},
	}))

	r.Get("/", adaptor.HTTPHandler(h.Handler()))
}
