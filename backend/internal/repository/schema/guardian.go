package schema

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type GuardianRepository struct {
	db *gorm.DB
}

func NewGuardianRepository(db *gorm.DB) *GuardianRepository {
	return &GuardianRepository{db: db}
}

// Create a new Guardian
func (r *GuardianRepository) Create(guardian *models.Guardian) error {
	return r.db.Create(guardian).Error
}

// Find all guardians of a patient
func (r *GuardianRepository) FindByPatient(patientID string) (*[]models.Guardian, error) {
	var guardians []models.Guardian
	subQuery := r.db.Model(&models.Patient{}).Select("guardian_id").Where("id = ?", patientID)

	err := r.db.Where("id = (?)", subQuery).Find(&guardians).Error
	return &guardians, err
}

func (r *GuardianRepository) FindByID(id string) (*models.Guardian, error) {
	var guardian models.Guardian
	err := r.db.Where("id = ?", id).First(&guardian).Error
	return &guardian, err
}

// Update guardian information
func (r *GuardianRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Guardian{}).Where("id = ?", id).Updates(updates).Error
}

// Delete a patient
func (r *GuardianRepository) Delete(id string) error {
	return r.db.Delete(&models.Guardian{}, "id = ?", id).Error
}
