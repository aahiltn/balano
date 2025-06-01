// internal/database/db.go

package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"palaam/internal/config"
	"palaam/internal/models" // Import models package
)

// seedDefaultAssessments creates default assessments if they don't exist
func seedDefaultAssessments(db *gorm.DB) error {
	defaultAssessments := []string{
		"VB-MAPP", // Verbal Behavior Milestones Assessment and Placement Program
		"ESFLS",   // Essential for Living Skills
		"ABLLS-R", // Assessment of Basic Language and Learning Skills-Revised
	}

	for _, name := range defaultAssessments {
		var count int64
		db.Model(&models.Assessment{}).Where("name = ?", name).Count(&count)

		if count == 0 {
			assessment := models.Assessment{
				Name: name,
			}
			if err := db.Create(&assessment).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// seedDefaultOnboardingQuestions creates default onboarding questions for each assessment
func seedDefaultOnboardingQuestions(db *gorm.DB) error {
	// Get all assessments
	var assessments []models.Assessment
	if err := db.Find(&assessments).Error; err != nil {
		return err
	}

	for _, assessment := range assessments {
		var defaultQuestions []models.OnboardingQuestion

		switch assessment.Name {
		case "VB-MAPP":
			defaultQuestions = []models.OnboardingQuestion{
				{Text: "Does the patient exhibit basic mand capabilities?", AssessmentID: assessment.ID},
				{Text: "Can the patient engage in spontaneous vocal behavior?", AssessmentID: assessment.ID},
				{Text: "Does the patient display listener responding skills?", AssessmentID: assessment.ID},
			}
		case "ESFLS":
			defaultQuestions = []models.OnboardingQuestion{
				{Text: "Can the patient make requests for essential items?", AssessmentID: assessment.ID},
				{Text: "Is the patient able to tolerate specific situations?", AssessmentID: assessment.ID},
				{Text: "Can the patient engage in daily living activities?", AssessmentID: assessment.ID},
			}
		case "ABLLS-R":
			defaultQuestions = []models.OnboardingQuestion{
				{Text: "How would you rate the patient's visual performance skills?", AssessmentID: assessment.ID},
				{Text: "Can the patient follow instructions?", AssessmentID: assessment.ID},
				{Text: "Does the patient demonstrate language comprehension?", AssessmentID: assessment.ID},
			}
		}

		for _, question := range defaultQuestions {
			var count int64
			db.Model(&models.OnboardingQuestion{}).Where("text = ?", question.Text).Count(&count)

			if count == 0 {
				if err := db.Create(&question).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func NewConnection(config *config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	// AutoMigrate to create/update tables based on models
	err = db.AutoMigrate(
		&models.Activity{},
		&models.Patient{},
		&models.Guardian{},
		&models.Staff{},
		&models.Medicine{},
		&models.OperatingHours{},
		&models.Branch{},
		&models.Session{},
		&models.Assessment{},
		&models.OnboardingQuestion{},
		&models.OnboardingResponse{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database schema: %w", err)
	}

	// Seed default assessments
	err = seedDefaultAssessments(db)
	if err != nil {
		return nil, fmt.Errorf("failed to seed default assessments: %w", err)
	}

	// Seed default onboarding questions for each assessment
	err = seedDefaultOnboardingQuestions(db)
	if err != nil {
		return nil, fmt.Errorf("failed to seed default onboarding questions: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
