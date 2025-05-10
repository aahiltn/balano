package schema

import (
	"gorm.io/gorm"

	"palaam/internal/models"
)

type BranchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{db: db} 
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

