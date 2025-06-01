package impl

// backend/internal/repository/impl/branch.go

import (
	"gorm.io/gorm"

	"palaam/internal/models"
)

type BranchRepository struct {
	db *gorm.DB
}

type OperatingHoursRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{db: db}
}

func NewOperatingHoursRepository(db *gorm.DB) *OperatingHoursRepository {
	return &OperatingHoursRepository{db: db}
}

func (r *BranchRepository) Create(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *BranchRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Branch{}).Where("id = ?", id).Updates(updates).Error
}

func (r *BranchRepository) GetBranchByID(id string) (*models.Branch, error) {
	var branch models.Branch
	err := r.db.Where("id = ?", id).First(&branch).Error
	return &branch, err
}

func (r *BranchRepository) ListBranches() (*models.Branch, error) {
	var branch models.Branch
	err := r.db.Find(&branch).Error
	return &branch, err
}

func (r *BranchRepository) DeleteBranch(id string) (*models.Branch, error) {
	var branch models.Branch
	err := r.db.Where("id = ?", id).First(&branch).Error
	return &branch, err
}

// Create new operating hours
func (r *OperatingHoursRepository) Create(hours *models.OperatingHours) error {
	return r.db.Create(hours).Error
}

// Find operating hours by BranchID and DayOfWeek
func (r *OperatingHoursRepository) FindByBranchAndDay(branchID int, dayOfWeek int16) (*models.OperatingHours, error) {
	var hours models.OperatingHours
	if err := r.db.First(&hours, "branch_id = ? AND day_of_week = ?", branchID, dayOfWeek).Error; err != nil {
		return nil, err
	}
	return &hours, nil
}

// Find all operating hours for a branch
func (r *OperatingHoursRepository) FindByBranch(branchID int) ([]*models.OperatingHours, error) {
	var hours []*models.OperatingHours
	if err := r.db.Where("branch_id = ?", branchID).Find(&hours).Error; err != nil {
		return nil, err
	}
	return hours, nil
}

// Update operating hours
func (r *OperatingHoursRepository) Update(branchID int, dayOfWeek int16, updates map[string]interface{}) error {
	return r.db.Model(&models.OperatingHours{}).
		Where("branch_id = ? AND day_of_week = ?", branchID, dayOfWeek).
		Updates(updates).Error
}

// Delete operating hours
func (r *OperatingHoursRepository) Delete(branchID int, dayOfWeek int16) error {
	return r.db.Where("branch_id = ? AND day_of_week = ?", branchID, dayOfWeek).
		Delete(&models.OperatingHours{}).Error
}
