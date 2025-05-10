package schema

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

// Create a new activity
func (r *ActivityRepository) Create(activity *models.Activity) error {
	return r.db.Create(activity).Error
}

// Find an activity by ID
func (r *ActivityRepository) FindByID(id string) (*models.Activity, error) {
	var activity models.Activity
	if err := r.db.First(&activity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// Find activities by SessionID
func (r *ActivityRepository) FindBySessionID(sessionID string) ([]*models.Activity, error) {
	var activities []*models.Activity
	if err := r.db.Where("session_id = ?", sessionID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

// Update an activity
func (r *ActivityRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Activity{}).Where("id = ?", id).Updates(updates).Error
}

// Delete an activity
func (r *ActivityRepository) Delete(id string) error {
	return r.db.Delete(&models.Activity{}, "id = ?", id).Error
}
