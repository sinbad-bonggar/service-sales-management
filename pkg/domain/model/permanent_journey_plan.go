package model

import (
	"time"
)

// PermanentJourneyPlan ...
type PermanentJourneyPlan struct {
	Base
	PortfolioID        uint            `json:"portfolioId" gorm:"index;not null;type:uint"`
	SupplierID         string          `json:"supplierId" gorm:"index;not null;type:string;size:100"`
	StoreID            string          `json:"storeId" gorm:"index;not null;type:string;size:100"`
	StoreExternalID    *string         `json:"storeExternalId" gorm:"index;type:string;size:256;default:null"`
	StoreName          string          `json:"storeName" gorm:"not null;type:string;size:256"`
	SalesRepID         string          `json:"salesRepId" gorm:"index;not null;type:string;size:100"`
	SalesRepExternalID *string         `json:"salesRepExternalId" gorm:"index;type:string;size:256;default:null"`
	SalesRepName       string          `json:"salesRepName" gorm:"not null;type:string;size:256"`
	IsRepetitive       bool            `json:"isRepetitive" gorm:"not null;type:bool;default:false"`
	RepeatEveryValue   *int            `json:"repeatEveryValue" gorm:"type:int;default:null"`
	RepeatEveryMeasure *string         `json:"repeatEveryMeasure" gorm:"default:null;type:string;size:100"`
	RepeatUntil        time.Time       `json:"repeatUntil" gorm:"not null"`
	Route              *int            `json:"route" gorm:"type:int;default:null"`
	VisitSchedules     *visitSchedules `json:"visitSchedules"`
}

type visitSchedules []PJPVisitSchedule

// VisitDateList ...
func (v visitSchedules) VisitDateList() []time.Time {
	var list []time.Time
	for _, visitSchedule := range v {
		list = append(list, visitSchedule.VisitDate)
	}

	return list
}
