package service

import (
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/model"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/repository"
)

type permanentJourneyPlanService struct {
	repository repository.PermanentJourneyPlanRepository
}

// PermanentJourneyPlanService ...
type PermanentJourneyPlanService interface {
	Get(id uint) (*model.PermanentJourneyPlan, error)
	GetAll(limit, offset int, search string) ([]*model.PermanentJourneyPlan, int64, error)
	Create(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error)
	Update(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error)
	Delete(id uint) error
}

// NewPermanentJourneyPlanService ...
func NewPermanentJourneyPlanService(r repository.PermanentJourneyPlanRepository) PermanentJourneyPlanService {
	return &permanentJourneyPlanService{
		repository: r,
	}
}

// GetPermanentJourneyPlan ...
func (s *permanentJourneyPlanService) Get(id uint) (*model.PermanentJourneyPlan, error) {

	return s.repository.Read(id)
}

// GetAllPermanentJourneyPlan ...
func (s *permanentJourneyPlanService) GetAll(limit, offset int, search string) ([]*model.PermanentJourneyPlan, int64, error) {

	return s.repository.ReadAll(limit, offset, search)
}

// CreatePermanentJourneyPlan ...
func (s *permanentJourneyPlanService) Create(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error) {

	return s.repository.Create(model)
}

// UpdatePermanentJourneyPlan ...
func (s *permanentJourneyPlanService) Update(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error) {

	return s.repository.Update(model)
}

// DeletePermanentJourneyPlan ...
func (s *permanentJourneyPlanService) Delete(id uint) error {

	return s.repository.Delete(id)
}
