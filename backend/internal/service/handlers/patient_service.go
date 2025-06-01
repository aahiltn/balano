package handlers

// handlers/patient_service.go
import (
	"errors"
	"palaam/internal/models"
	"palaam/internal/repository"
)

// PatientService implements PatientServiceInterface
type PatientService struct {
	repo *repository.Repository
}

// NewPatientService creates a new PatientService
func NewPatientService(repo *repository.Repository) PatientServiceInterface {
	return &PatientService{
		repo: repo,
	}
}

// Create creates a new patient
func (s *PatientService) Create(patient *models.Patient) (*models.Patient, error) {
	// Business logic validation
	if patient.Name == "" {
		return nil, errors.New("patient name is required")
	}

	// Additional validation based on your Patient model
	// e.g., check for duplicate patients, validate fields, etc.

	return s.repo.Patient.Create(patient)
}

// GetByID retrieves a patient by ID
func (s *PatientService) GetByID(id uint) (*models.Patient, error) {
	if id == 0 {
		return nil, errors.New("invalid patient ID")
	}

	patient, err := s.repo.Patient.GetByID(id)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, errors.New("patient not found")
	}

	return patient, nil
}

// Update updates a patient
func (s *PatientService) Update(id uint, patient *models.Patient) (*models.Patient, error) {
	// Check if patient exists
	_, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Business logic for updates
	if patient.Name == "" {
		return nil, errors.New("patient name cannot be empty")
	}

	return s.repo.Patient.Update(id, patient)
}

// Delete deletes a patient
func (s *PatientService) Delete(id uint) error {
	// Check if patient exists
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	// Additional business logic - check if patient has active sessions
	sessions, _, err := s.GetSessions(id, 1, 0)
	if err != nil {
		return err
	}

	if len(sessions) > 0 {
		return errors.New("cannot delete patient with existing sessions")
	}

	return s.repo.Patient.Delete(id)
}

// List retrieves a list of patients with pagination
func (s *PatientService) List(limit, offset int) ([]*models.Patient, int64, error) {
	// Validate pagination parameters
	if limit <= 0 {
		limit = 10 // default limit
	}
	if limit > 100 {
		limit = 100 // max limit
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.Patient.List(limit, offset)
}

// GetSessions retrieves all sessions for a patient
func (s *PatientService) GetSessions(patientID uint, limit, offset int) ([]*models.Session, int64, error) {
	// Check if patient exists
	_, err := s.GetByID(patientID)
	if err != nil {
		return nil, 0, err
	}

	// Validate pagination parameters
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.Session.GetByPatientID(patientID, limit, offset)
}
