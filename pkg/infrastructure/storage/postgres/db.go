package postgres

import (
	"fmt"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/config"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/utilities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDBConnection represent database config
func NewDBConnection() (*gorm.DB, *gorm.DB) {
	Config := config.NewConfig()

	// connecting to sales management
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Config.DBHost,
		Config.DBUsername,
		Config.DBPassword,
		Config.DBName,
		Config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	utilities.Log("Db connection %s SUCCESS", Config.DBName)

	// // connecting to medea
	// dsn = fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	Config.DBMedeaHost,
	// 	Config.DBMedeaUsername,
	// 	Config.DBMedeaPassword,
	// 	Config.DBMedeaName,
	// 	Config.DBMedeaPort,
	// )

	// dbMedea, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// sqlDbMedea, err := dbMedea.DB()
	// if err != nil {
	// 	panic(err)
	// }

	// err = sqlDbMedea.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// utilities.Log("Db connection %s SUCCESS", Config.DBMedeaName)

	return db, nil
}
