package impl

// backend/internal/repository/impl/onboarding_response.go

import (
	"palaam/internal/models"
	"time"

	"gorm.io/gorm"
)

type OnboardingResponseRepository struct {
	db *gorm.DB
}

func NewOnboardingResponseRepository(db *gorm.DB) *OnboardingResponseRepository {
	return &OnboardingResponseRepository{db: db}
}

// Create a new onboarding response
func (r *OnboardingResponseRepository) Create(response *models.OnboardingResponse) error {
	if response.ResponseDate.IsZero() {
		response.ResponseDate = time.Now()
	}
	return r.db.Create(response).Error
}

// Find an onboarding response by ID
func (r *OnboardingResponseRepository) FindByID(id int) (*models.OnboardingResponse, error) {
	var response models.OnboardingResponse
	if err := r.db.First(&response, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &response, nil
}

// Find onboarding responses by patient ID
func (r *OnboardingResponseRepository) FindByPatientID(patientID string) ([]*models.OnboardingResponse, error) {
	var responses []*models.OnboardingResponse
	if err := r.db.Where("patient_id = ?", patientID).Find(&responses).Error; err != nil {
		return nil, err
	}
	return responses, nil
}

// Find onboarding responses by question text
func (r *OnboardingResponseRepository) FindByPatientAndQuestion(patientID string, questionText string) ([]*models.OnboardingResponse, error) {
	var responses []*models.OnboardingResponse
	if err := r.db.Where("patient_id = ? AND question_text = ?", patientID, questionText).Find(&responses).Error; err != nil {
		return nil, err
	}
	return responses, nil
}

// Update an onboarding response
func (r *OnboardingResponseRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&models.OnboardingResponse{}).Where("id = ?", id).Updates(updates).Error
}

// Delete an onboarding response
func (r *OnboardingResponseRepository) Delete(id int) error {
	return r.db.Delete(&models.OnboardingResponse{}, "id = ?", id).Error
}

// CreateInitialOnboardingResponses creates initial placeholder responses for a new patient
// This is useful when a new patient is registered and needs to go through the onboarding process
func (r *OnboardingResponseRepository) CreateInitialOnboardingResponses(patientID string, staffID string, assessmentID int) error {
	// Get all questions for the selected assessment
	var questions []*models.OnboardingQuestion
	if err := r.db.Where("assessment_id = ?", assessmentID).Find(&questions).Error; err != nil {
		return err
	}

	// Create a response entry for each question
	for _, question := range questions {
		response := models.OnboardingResponse{
			QuestionText: question.Text,
			PatientID:    patientID,
			StaffID:      staffID,
			ResponseDate: time.Now(),
		}

		if err := r.db.Create(&response).Error; err != nil {
			return err
		}
	}

	return nil
}
