package impl

// backend/internal/repository/impl/assessments.go

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type AssessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository(db *gorm.DB) *AssessmentRepository {
	return &AssessmentRepository{db: db}
}

// Create a new assessment
func (r *AssessmentRepository) Create(assessment *models.Assessment) error {
	return r.db.Create(assessment).Error
}

// Find an assessment by ID
func (r *AssessmentRepository) FindByID(id int) (*models.Assessment, error) {
	var assessment models.Assessment
	if err := r.db.First(&assessment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

// Find assessment by name
func (r *AssessmentRepository) FindByName(name string) (*models.Assessment, error) {
	var assessment models.Assessment
	if err := r.db.Where("name = ?", name).First(&assessment).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

// Get all assessments
func (r *AssessmentRepository) GetAll() ([]*models.Assessment, error) {
	var assessments []*models.Assessment
	if err := r.db.Find(&assessments).Error; err != nil {
		return nil, err
	}
	return assessments, nil
}

// Update an assessment
func (r *AssessmentRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&models.Assessment{}).Where("id = ?", id).Updates(updates).Error
}

// Delete an assessment
func (r *AssessmentRepository) Delete(id int) error {
	return r.db.Delete(&models.Assessment{}, "id = ?", id).Error
}

// EnsureDefaultAssessments creates default assessments if they don't exist
func (r *AssessmentRepository) EnsureDefaultAssessments() error {
	defaultAssessments := []string{
		"VB-MAPP", // Verbal Behavior Milestones Assessment and Placement Program
		"ESFLS",   // Essential for Living Skills
		"ABLLS-R", // Assessment of Basic Language and Learning Skills-Revised
	}

	for _, name := range defaultAssessments {
		var count int64
		r.db.Model(&models.Assessment{}).Where("name = ?", name).Count(&count)

		if count == 0 {
			assessment := models.Assessment{
				Name: name,
			}
			if err := r.db.Create(&assessment).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
