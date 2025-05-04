// internal/models/models.go
package models

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID              string          `db:"id"`
	Description     sql.NullString  `db:"description"`
	DurationMinutes sql.NullFloat64 `db:"duration_minutes"`
	PaymentReceived sql.NullBool    `db:"payment_received"`
	SessionID       sql.NullString  `db:"session_id"`
	CreatedAt       time.Time       `db:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at"`
}

type Patient struct {
	ID           string         `db:"id"`
	Name         string         `db:"name"`
	Dob          string         `db:"dob"`
	Active       sql.NullBool   `db:"active"`
	GuardianID   sql.NullString `db:"guardian_id"`
	DoctorID     sql.NullString `db:"doctor_id"`
	StaffID      string         `db:"staff_id"`
	TherapyTypes sql.NullString `db:"therapy_types"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

type Guardian struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Email       sql.NullString `db:"email"`
}

type StaffRole string

type Staff struct {
	ID            string    `db:"id"`
	Name          string    `db:"name"`
	JoinDate      time.Time `db:"join_date"`
	ExpectedHours int       `db:"expected_hours"`
	Role          StaffRole `db:"role"`
}

type Medicines struct {
	ID           string         `db:"id"`
	Name         string         `db:"name"`
	BrandName    sql.NullString `db:"brand_name"`
	PatientID    string         `db:"patient_id"`
	PrescriberID string         `db:"prescriber_id"`
}

type OperatingHours struct {
	DayOfWeek int    `db:"day_of_week"` // 0 for Sunday, 1 for Monday, etc.
	OpenTime  string `db:"open_time"`   // Store as "HH:MM" format
	CloseTime string `db:"close_time"`  // Store as "HH:MM" format
	IsClosed  bool   `db:"is_closed"`   // If the branch is closed on this day
	BranchID  int16  `db:"branch_id"`   // Foreign key to Branch
}

type Branch struct {
	ID          int            `db:"id"`
	Location    sql.NullString `db:"location"`
	OpeningDate time.Time      `db:"opening_date"`
	Active      bool           `db:"active"`
}

type ResponseLevel string

type Session struct {
	ID              string        `db:"id"`
	PatientID       string        `db:"patient_id"`
	StaffID         string        `db:"patient_id"`
	StartTime       time.Time     `db:"start_time"`
	EndTime         time.Time     `db:"end_time"`
	Description     string        `db:"description"`
	Response        ResponseLevel `db:"response"`
	PaymentReceived sql.NullBool  `db:"payment_received"`
}
