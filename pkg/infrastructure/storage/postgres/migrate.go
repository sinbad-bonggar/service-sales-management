package postgres

import (
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/model"
	"gorm.io/gorm"
)

// Migrate represent migration schema models
func Migrate(db *gorm.DB) error {
	PermanentJourneyPlan := model.PermanentJourneyPlan{}
	PJPVisitSchedule := model.PJPVisitSchedule{}

	db.Exec("DO $$ BEGIN CREATE TYPE  repeatoption AS ENUM('week', 'month'); EXCEPTION WHEN DUPLICATE_OBJECT THEN null; END $$;")

	err := db.AutoMigrate(&PermanentJourneyPlan, &PJPVisitSchedule)

	return err
}
