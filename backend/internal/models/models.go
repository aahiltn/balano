// internal/models/models.go
package models

import (
	"time"
)

type Activity struct {
	ID              string  `gorm:"primaryKey;type:uuid"`
	Description     *string `gorm:"type:text"`
	DurationMinutes *float64
	SessionID       *string `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ResponseLevel   *ResponseLevel `gorm:"type:text"`

	// Relationships
	Session *Session `gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Patient struct {
	ID              string `gorm:"primaryKey;type:uuid"`
	Name            string
	Dob             string
	Active          *bool
	DoctorID        *string `gorm:"type:uuid"`
	PrimaryBranchID *int    `gorm:"type:int"`
	TherapyTypes    *string
	JoinDate        time.Time
	UpdatedAt       time.Time

	// Relationships
	Guardians []*Guardian `gorm:"many2many:patient_guardians;"`
	Doctor    *Staff      `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Branch    Branch      `gorm:"foreignKey:PrimaryBranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Guardian struct {
	ID          string `gorm:"primaryKey;type:uuid"`
	Name        string
	PhoneNumber *string
	Email       *string

	// Relationships
	Patients []*Patient `gorm:"many2many:patient_guardians;"`
}

type StaffRole string

type Staff struct {
	ID              string `gorm:"primaryKey;type:uuid"`
	Name            string
	JoinDate        time.Time
	ExpectedHours   int
	Role            StaffRole `gorm:"type:varchar(50)"`
	PrimaryBranchID *int      `gorm:"type:int"`

	// Relationships
	Patients  []Patient  `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Sessions  []Session  `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Medicines []Medicine `gorm:"foreignKey:PrescriberID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Branch    Branch     `gorm:"foreignKey:PrimaryBranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Medicine struct {
	ID           string `gorm:"primaryKey;type:uuid"`
	Name         string
	BrandName    *string
	PatientID    string `gorm:"type:uuid"`
	PrescriberID string `gorm:"type:uuid"` // Refers to Staff (Doctor)

	// Relationships
	Patient    Patient `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Prescriber Staff   `gorm:"foreignKey:PrescriberID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type OperatingHours struct {
	DayOfWeek int16  `gorm:"primaryKey"` // 0 for Sunday, etc.
	OpenTime  string `gorm:"type:varchar(5)"`
	CloseTime string `gorm:"type:varchar(5)"`
	IsClosed  bool
	BranchID  int `gorm:"primaryKey"` // Composite PK

	// Relationships
	Branch Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Branch struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Location    *string
	OpeningDate time.Time
	Active      bool

	// Relationships
	OperatingHours []OperatingHours `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ResponseLevel string // low medium high

type Session struct {
	ID              string `gorm:"primaryKey;type:uuid"`
	PatientID       string `gorm:"type:uuid"`
	StaffID         string `gorm:"type:uuid"`
	StartTime       time.Time
	EndTime         time.Time
	Description     string
	Response        ResponseLevel `gorm:"type:varchar(50)"`
	PaymentReceived *bool

	// Relationships
	Patient    Patient    `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Staff      Staff      `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Activities []Activity `gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Assessment struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(100);not null;unique"`
}

type OnboardingQuestion struct {
	Text         string `gorm:"primaryKey;type:varchar(100);not null;unique"`
	Group        int
	AssessmentID int

	Assessment Assessment `gorm:"foreignKey:AssessmentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type OnboardingResponse struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	QuestionText string
	PatientID    string    `gorm:"type:uuid"`
	StaffID      string    `gorm:"type:uuid"`
	SessionID    *string   `gorm:"type:uuid"`
	ResponseDate time.Time `gorm:"autoCreateTime"`

	Patient            Patient            `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Staff              Staff              `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	OnboardingQuestion OnboardingQuestion `gorm:"foreignKey:QuestionText;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Session            *Session           `gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
