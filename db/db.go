package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"

	"go-vote/config"
)

func LoadDb(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Errorf("failed to connect to database: %v", err)
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(cfg.DbConnMaxLifetimeMinute))
	db.SetMaxOpenConns(cfg.DbMaxOpenIdleConn)
	db.SetMaxIdleConns(cfg.DbMaxOpenIdleConn)
	log.Infof("database connected to %s", cfg.DbDriver)
	return db, nil
}
