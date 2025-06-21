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

// NewSessionRepository creates a new instance of SessionRepository
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
func (r *SessionRepository) FindByDateRange(branchId int, startDate, endDate time.Time) ([]*models.Session, error) {
	var sessions []*models.Session
	if err := r.db.Where("branch_id = ? AND start_time >= ? AND end_time <= ?", branchId, startDate, endDate).Find(&sessions).Error; err != nil {
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
func (r *SessionRepository) Update(id string, updates map[string]interface{}) (*models.Session, error) {
	var session models.Session
	if err := r.db.Model(&session).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// Delete a session
func (r *SessionRepository) Delete(id string) error {
	return r.db.Delete(&models.Session{}, "id = ?", id).Error
}

// CheckOverlappingSessions checks if there are any sessions overlapping with the given time range for a specific staff
func (r *SessionRepository) CheckOverlappingSessions(staffID string, startTime, endTime time.Time, excludeSessionID string) (bool, error) {
	var count int64
	query := r.db.Model(&models.Session{}).
		Where("staff_id = ? AND ((start_time < ? AND end_time > ?) OR (start_time >= ? AND start_time < ?))",
			staffID, endTime, startTime, startTime, endTime)

	// Exclude the current session when checking for updates
	if excludeSessionID != "" {
		query = query.Where("id <> ?", excludeSessionID)
	}

	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
