package impl

// backend/internal/repository/impl/patient.go

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

// Create a new patient
func (r *PatientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

// Find a patient by ID
func (r *PatientRepository) FindByID(id string) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.First(&patient, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

// FindByName finds patients by name (partial match)
func (r *PatientRepository) FindByName(name string) ([]*models.Patient, error) {
	var patients []*models.Patient
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

// Update a patient
func (r *PatientRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Patient{}).Where("id = ?", id).Updates(updates).Error
}

// Delete a patient
func (r *PatientRepository) Delete(id string) error {
	return r.db.Delete(&models.Patient{}, "id = ?", id).Error
}
