package main

import (
	"go.uber.org/fx"

	"github.com/stephane-nguyen/FamilyTree/server/config"
	"github.com/stephane-nguyen/FamilyTree/server/database"
	"github.com/stephane-nguyen/FamilyTree/server/handler"
	"github.com/stephane-nguyen/FamilyTree/server/router"
	"github.com/stephane-nguyen/FamilyTree/server/storage"
)

func main() {
	fx.New(
		fx.Provide(
			config.LoadEnv,
			database.CreateMySqlConnection,
			handler.NewPersonHandler,
			storage.NewPersonStorage,
		),
		fx.Invoke(router.NewFiberServer),
	).Run()
}