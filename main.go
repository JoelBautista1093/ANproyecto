package main

import (
	"context"

	"github.com/JoelBautista1093/ANproyecto/database"
	"github.com/JoelBautista1093/ANproyecto/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
