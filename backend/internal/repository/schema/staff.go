package schema

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db: db}
}

// Create a new staff member
func (r *StaffRepository) Create(staff *models.Staff) error {
	return r.db.Create(staff).Error
}

// Find a staff member by ID
func (r *StaffRepository) FindByID(id string) (*models.Staff, error) {
	var staff models.Staff
	if err := r.db.First(&staff, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

// Find staff members by role
func (r *StaffRepository) FindByRole(role models.StaffRole) ([]*models.Staff, error) {
	var staff []*models.Staff
	if err := r.db.Where("role = ?", role).Find(&staff).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

// Update a staff member
func (r *StaffRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Staff{}).Where("id = ?", id).Updates(updates).Error
}

// Delete a staff member
func (r *StaffRepository) Delete(id string) error {
	return r.db.Delete(&models.Staff{}, "id = ?", id).Error
}
