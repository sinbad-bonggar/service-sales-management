package repository

import "github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/model"

// PermanentJourneyPlanRepository represent repository of the permanent journey plan
type PermanentJourneyPlanRepository interface {
	Create(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error)
	Read(id uint) (*model.PermanentJourneyPlan, error)
	ReadAll(limit, offset int, search string) ([]*model.PermanentJourneyPlan, int64, error)
	Update(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error)
	Delete(id uint) error
}
