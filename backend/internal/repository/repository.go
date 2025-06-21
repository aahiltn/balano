package repository

import (
	"palaam/internal/repository/impl"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Session:  impl.NewSessionRepository(db),
		Activity: impl.NewActivityRepository(db),
		Patient:  impl.NewPatientRepository(db),
		Staff:    impl.NewStaffRepository(db),
	}
}
