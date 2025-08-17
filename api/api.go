package api

import (
	"context"
	"fmt"

	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
	"github.com/nitinjangam/pos-receipt-system/internal/config"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"golang.org/x/sync/errgroup"
)

type Services struct {
	ProductService  service.ProductServiceInterface
	SalesService    service.SalesServiceInterface
	AuthService     service.AuthServiceInterface
	SettingsService service.SettingsServiceInterface
}

func Run(ctx context.Context, c config.Config, handler v1.ServerInterface) error {
	g, ctx := errgroup.WithContext(ctx)

	// API v1
	g.Go(func() error {
		apiHandler, _ := v1.New(&v1.Config{
			Port:     fmt.Sprintf(":%s", c.Port),
			Services: []v1.ServerInterface{handler},
			Logger:   c.Logger,
		})
		if err := apiHandler.Run(ctx); err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
