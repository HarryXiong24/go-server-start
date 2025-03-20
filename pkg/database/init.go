package database

import (
	"go-server-start/internal/config"
	"go-server-start/pkg/logger"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Database instance
var (
	db   *gorm.DB
	once sync.Once
)

// Database options
type Options struct {
	DSN        string       // Database connection string
	GormConfig *gorm.Config // Gorm configuration
	DriverName string       // Database driver name
}

// DefaultOptions returns default database options
func DefaultOptions() *Options {
	return &Options{
		DSN:        config.AppConfig.Database.DBName,
		DriverName: config.AppConfig.Database.Driver,
		GormConfig: &gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // Use singular table names
			},
			Logger: logger.NewGormLogger(),
		},
	}
}

// Init initializes database connection
func Init() error {
	options := DefaultOptions()
	return InitWithOptions(options)
}

// InitWithOptions initializes database with custom options
func InitWithOptions(options *Options) error {
	var err error

	// Use sync.Once to ensure initialization happens only once
	once.Do(func() {
		logger.Sugar.Infof("Initializing database with DSN: %s", options.DSN)

		switch options.DriverName {
		case "sqlite3":
			db, err = gorm.Open(sqlite.Open(options.DSN), options.GormConfig)
		// Add other database drivers here, such as MySQL, PostgreSQL, etc.
		default:
			// Default to sqlite
			db, err = gorm.Open(sqlite.Open(options.DSN), options.GormConfig)
		}

		if err != nil {
			logger.Sugar.Errorf("Failed to connect to database: %v", err)
			return
		}

		logger.Sugar.Info("Database initialized successfully")
	})

	return err
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	if db == nil {
		logger.Sugar.Error("Database not initialized, please call Init() first")
		// Try to initialize automatically
		if err := Init(); err != nil {
			panic("Unable to auto-initialize database: " + err.Error())
		}
	}
	return db
}

// SetDB sets the global database instance (mainly for testing)
func SetDB(newDB *gorm.DB) {
	db = newDB
}

// Close closes the database connection
func Close() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// WithTransaction executes a function within a transaction
func WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := GetDB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
