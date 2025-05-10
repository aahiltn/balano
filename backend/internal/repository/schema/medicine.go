package schema

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type MedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}

// Create a new medicine
func (r *MedicineRepository) Create(medicine *models.Medicine) error {
	return r.db.Create(medicine).Error
}

// Find a medicine by ID
func (r *MedicineRepository) FindByID(id string) (*models.Medicine, error) {
	var medicine models.Medicine
	if err := r.db.First(&medicine, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &medicine, nil
}

// Find medicines by PatientID
func (r *MedicineRepository) FindByPatientID(patientID string) ([]*models.Medicine, error) {
	var medicines []*models.Medicine
	if err := r.db.Where("patient_id = ?", patientID).Find(&medicines).Error; err != nil {
		return nil, err
	}
	return medicines, nil
}

// Find medicines by PrescriberID
func (r *MedicineRepository) FindByPrescriberID(prescriberID string) ([]*models.Medicine, error) {
	var medicines []*models.Medicine
	if err := r.db.Where("prescriber_id = ?", prescriberID).Find(&medicines).Error; err != nil {
		return nil, err
	}
	return medicines, nil
}

// Update a medicine
func (r *MedicineRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Medicine{}).Where("id = ?", id).Updates(updates).Error
}

// Delete a medicine
func (r *MedicineRepository) Delete(id string) error {
	return r.db.Delete(&models.Medicine{}, "id = ?", id).Error
}
