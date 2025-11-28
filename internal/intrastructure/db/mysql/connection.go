package mysql

import (
	"github.com/Mohamadreza-shad/auth/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.Database) (*gorm.DB, error) {
	gormCfg := gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gormCfg)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(1 * time.Minute)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	return db, nil
}
