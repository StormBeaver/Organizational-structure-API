package database

import (
	"fmt"
	"orgService/internal/config"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// init PostgreSQL
func ConfigureGorm(cfg *config.Config, logger *zerolog.Logger) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		DisableAutomaticPing: true,
	}

	conn := postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v connect_timeout=10",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.SslMode,
		),
	})

	// SQL logger
	if cfg.Project.Debug {
		gormConfig.Logger = gormLogger.New(
			logger,
			gormLogger.Config{
				SlowThreshold:             20 * time.Millisecond,
				LogLevel:                  gormLogger.Info,
				IgnoreRecordNotFoundError: true,
			},
		)
	}

	db, err := gorm.Open(conn, gormConfig)
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sql.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
