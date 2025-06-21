package repository

import (
	"palaam/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	Assessment         AssessmentRepository
	OperatingHours     OperatingHoursRepository
	Staff              StaffRepository
	Activity           ActivityRepository
	Session            SessionRepository
	Patient            PatientRepository
	Guardian           GuardianRepository
	OnboardingQuestion OnboardingQuestionRepository
	OnboardingResponse OnboardingResponseRepository
	Medicine           MedicineRepository
	Branch             BranchRepository
}

// AssessmentRepository defines the interface for assessment repository operations
type AssessmentRepository interface {
	NewAssessmentRepository(db *gorm.DB) *AssessmentRepository
	Create(assessment *models.Assessment) error
	FindByID(id int) (*models.Assessment, error)
	FindByName(name string) (*models.Assessment, error)
	GetAll() ([]*models.Assessment, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
	EnsureDefaultAssessments() error
}

type OperatingHoursRepository interface {
	NewOperatingHoursRepository(db *gorm.DB) *OperatingHoursRepository
	Create(hours *models.OperatingHours) error
	FindByBranchAndDay(branchID int, dayOfWeek int16) (*models.OperatingHours, error)
	FindByBranch(branchID int) ([]*models.OperatingHours, error)
	Update(branchID int, dayOfWeek int16, updates map[string]interface{}) error
	Delete(branchID int, dayOfWeek int16) error
}

type StaffRepository interface {
	NewStaffRepository(db *gorm.DB) *StaffRepository
	Create(staff *models.Staff) error
	FindByID(id string) (*models.Staff, error)
	FindByRole(role models.StaffRole) ([]*models.Staff, error)
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}

type ActivityRepository interface {
	NewActivityRepository(db *gorm.DB) *ActivityRepository
	Create(activity *models.Activity) error
	FindByID(id string) (*models.Activity, error)
	FindBySessionID(name string) ([]*models.Activity, error)
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}

type SessionRepository interface {
	Create(session *models.Session) error
	FindByID(id string) (*models.Session, error)
	FindByDateRange(branchID int, startDate, endDate time.Time) ([]*models.Session, error)
	FindByPatientID(patientID string) ([]*models.Session, error)
	FindByStaffID(staffID string) ([]*models.Session, error)
	Update(id string, updates map[string]interface{}) (*models.Session, error)
	Delete(id string) error
	CheckOverlappingSessions(patientID string, startTime, endTime time.Time, excludeSessionId string) (bool, error)
}

type PatientRepository interface {
	NewPatientRepository(db *gorm.DB) *PatientRepository
	Create(patient *models.Patient) error
	FindByID(id string) (*models.Patient, error)
	FindByName(name string) (*models.Patient, error)
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}

type GuardianRepository interface {
	NewGuardianRepository(db *gorm.DB) *GuardianRepository
	Create(guardian *models.Guardian) error
	FindByPatient(patientID string) (*[]models.Guardian, error)
	FindByID(id string) (*models.Guardian, error)
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}

type OnboardingQuestionRepository interface {
	NewOnboardingQuestionRepository(db *gorm.DB) *OnboardingQuestionRepository
	Create(question *models.OnboardingQuestion) error
	FindByText(text string) (*models.OnboardingQuestion, error)
	FindByAssessmentID(assessmentID int) ([]*models.OnboardingQuestion, error)
	GetAll() ([]*models.OnboardingQuestion, error)
	Update(text string, updates map[string]interface{}) error
	Delete(text string) error
	CreateDefaultQuestionsForAssessment(assessmentID int, assessmentName string) error
}

type OnboardingResponseRepository interface {
	NewOnboardingResponseRepository(db *gorm.DB) *OnboardingResponseRepository
	Create(response *models.OnboardingResponse) error
	FindByID(id int) (*models.OnboardingResponse, error)
	FindByPatientID(patientID string) ([]*models.OnboardingResponse, error)
	FindByPatientAndQuestion(patientID string, questionText string) ([]*models.OnboardingResponse, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
}

type MedicineRepository interface {
	NewMedicineRepository(db *gorm.DB) *MedicineRepository
	Create(medicine *models.Medicine) error
	FindByID(id int) (*models.Medicine, error)
	FindByPatientID(patientID string) ([]*models.Medicine, error)
	FindByPrescriberID(prescriberID string) ([]*models.Medicine, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
}

type BranchRepository interface {
	NewBranchRepository(db *gorm.DB) *BranchRepository
	Create(branch *models.Branch) error
	Update(id string, updates map[string]interface{}) error
	GetBranchByID(id string) (*models.Branch, error)
	ListBranches() ([]*models.Branch, error)
	DeleteBranch(id string) (*models.Branch, error)
}
