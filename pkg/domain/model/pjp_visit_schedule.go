package model

import "time"

// PJPVisitSchedule ...
type PJPVisitSchedule struct {
	Base
	PermanentJourneyPlanID uint      `json:"permanentJourneyPlanID" gorm:"index;column:permanent_journey_plan_id;not null;type:uint"`
	VisitDate              time.Time `json:"visitDate" gorm:"not null;type:date"`
}
