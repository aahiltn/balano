package impl

// backend/internal/repository/impl/session.go

import (
	"time"

	"palaam/internal/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// Create a new session
func (r *SessionRepository) Create(session *models.Session) error {
	return r.db.Create(session).Error
}

// Find a session by ID
func (r *SessionRepository) FindByID(id string) (*models.Session, error) {
	var session models.Session
	if err := r.db.First(&session, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// Find sessions by date range
func (r *SessionRepository) FindByDateRange(startDate, endDate time.Time) ([]*models.Session, error) {
	var sessions []*models.Session
	if err := r.db.Where("start_time >= ? AND end_time <= ?", startDate, endDate).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// Find sessions by PatientID
func (r *SessionRepository) FindByPatientID(patientID string) ([]*models.Session, error) {
	var sessions []*models.Session
	if err := r.db.Where("patient_id = ?", patientID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// Find sessions by StaffID
func (r *SessionRepository) FindByStaffID(staffID string) ([]*models.Session, error) {
	var sessions []*models.Session
	if err := r.db.Where("staff_id = ?", staffID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// Update a session
func (r *SessionRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Session{}).Where("id = ?", id).Updates(updates).Error
}

// Delete a session
func (r *SessionRepository) Delete(id string) error {
	return r.db.Delete(&models.Session{}, "id = ?", id).Error
}
