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
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database schema: %w", err)
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