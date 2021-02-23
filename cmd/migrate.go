package main

import (
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/utilities"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/storage/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, _ := postgres.NewDBConnection()

	utilities.Log("DB connection %s SUCCESS", db.Name())

	err := postgres.Migrate(db)
	if err != nil {
		panic(err)
	}

	utilities.Log("Success migration")
}
