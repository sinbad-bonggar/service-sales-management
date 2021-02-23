package response

import (
	"time"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/model"
)

// PermanentJourneyPlanResponse ...
type PermanentJourneyPlanResponse struct {
	ID                 uint        `json:"id"`
	CreatedAt          time.Time   `json:"createdAt"`
	UpdatedAt          time.Time   `json:"updatedAt"`
	DeletedAt          *time.Time  `json:"deletedAt,omitempty"`
	PortfolioID        uint        `json:"PortfolioID"`
	SupplierID         string      `json:"SupplierID"`
	StoreID            string      `json:"storeID"`
	StoreExternalID    *string     `json:"storeExternalID"`
	StoreName          string      `json:"storeName"`
	SalesRepID         string      `json:"salesRepID"`
	SalesRepExternalID *string     `json:"salesRepExternalID"`
	SalesRepName       string      `json:"salesRepName"`
	IsRepetitive       bool        `json:"isRepetitive"`
	RepeatEveryValue   *int        `json:"repeatEveryValue"`
	RepeatEveryMeasure *string     `json:"repeatEveryMeasure"`
	RepeatUntil        time.Time   `json:"repeatUntil"`
	Route              *int        `json:"route"`
	VisitSchedules     []time.Time `json:"visitSchedules"`
}

// MapFromModel ...
func MapFromModel(model *model.PermanentJourneyPlan) PermanentJourneyPlanResponse {
	res := PermanentJourneyPlanResponse{}
	res.PortfolioID = model.PortfolioID
	res.StoreID = model.StoreID
	res.StoreExternalID = model.StoreExternalID
	res.StoreName = model.StoreName
	res.SalesRepID = model.SalesRepID
	res.SalesRepExternalID = model.SalesRepExternalID
	res.SalesRepName = model.SalesRepName
	res.IsRepetitive = model.IsRepetitive
	res.RepeatEveryValue = model.RepeatEveryValue
	res.RepeatEveryMeasure = model.RepeatEveryMeasure
	res.RepeatUntil = model.RepeatUntil
	res.Route = model.Route
	res.VisitSchedules = model.VisitSchedules.VisitDateList()

	return res
}
