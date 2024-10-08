package main

import (
	"context"

	"github.com/JoelBautista1093/ANproyecto/database"
	"github.com/JoelBautista1093/ANproyecto/internal/repository"
	"github.com/JoelBautista1093/ANproyecto/internal/service"
	"github.com/JoelBautista1093/ANproyecto/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
				if err != nil {
					panic(err)
				}
				if u.Name != "myname" {
					panic("wrong name")
				}
			},
		),
	)
	app.Run()
}
