package postgres

import (
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/model"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/domain/repository"
	"gorm.io/gorm"
)

// permanentJourneyPlanRepository implements repository.PermanentJourneyPlanRepository
type permanentJourneyPlanRepository struct {
	db *gorm.DB
}

// NewPermanentJourneyPlanRepository returns initialized PermanentJourneyPlanRepository
func NewPermanentJourneyPlanRepository(db *gorm.DB) repository.PermanentJourneyPlanRepository {
	return &permanentJourneyPlanRepository{
		db: db,
	}
}

func (repo *permanentJourneyPlanRepository) Create(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error) {

	result := repo.db.Create(&model)

	err := result.Scan(&model).Error

	if err != nil {
		return model, err
	}

	return model, err
}

func (repo *permanentJourneyPlanRepository) Read(id uint) (*model.PermanentJourneyPlan, error) {
	model := model.PermanentJourneyPlan{}

	err := repo.db.First(&model, id).Error

	return &model, err
}

func (repo *permanentJourneyPlanRepository) ReadAll(limit, offset int, search string) ([]*model.PermanentJourneyPlan, int64, error) {
	model := []*model.PermanentJourneyPlan{}

	var count int64

	err := repo.db.Find(&model).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = repo.db.Limit(limit).Offset(offset).Order("created_at DESC").Preload("VisitSchedules").Find(&model).Error

	return model, count, err
}

func (repo *permanentJourneyPlanRepository) Update(model *model.PermanentJourneyPlan) (*model.PermanentJourneyPlan, error) {
	result := repo.db.Save(&model)
	err := result.Error

	return model, err
}

func (repo *permanentJourneyPlanRepository) Delete(id uint) error {
	model := model.PermanentJourneyPlan{}

	result := repo.db.Model(&model).Delete("id = ?", id)

	return result.Error
}
