package infra

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"

	"go-vote/config"
	sql2 "go-vote/infra/db"
	redis2 "go-vote/infra/redis"
)

type Infra struct {
	Config *config.Config
	Db     *sql.DB
	Redis  *redis.Client
}

func Init() (*Infra, error) {
	// load config
	cfg, err := config.Init()
	if err != nil {
		log.Errorf("failed to load config: %v", err)
		return nil, err
	}

	// load db
	db, err := sql2.Init(*cfg)
	if err != nil {
		log.Errorf("failed to load db: %v", err)
		return nil, err
	}

	// load redis
	rd := redis2.Init(*cfg)

	inf := &Infra{
		Config: cfg,
		Db:     db,
		Redis:  rd,
	}

	return inf, nil
}
