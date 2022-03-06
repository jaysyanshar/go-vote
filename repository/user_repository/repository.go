package user_repository

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"go-vote/infra"
	"go-vote/infra/db"
	redis2 "go-vote/infra/redis"

	"github.com/labstack/gommon/log"

	"go-vote/model"
)

type UserRepository interface {
	Insert(user *model.InsertUserDb) (int64, error)
	FindByEmail(email string) (*model.FindUserDb, error)
	FindById(id int64) (*model.FindUserDb, error)
}

type repository struct {
	Db    *sql.DB
	Redis *redis.Client
}

const (
	userCacheExpirationSecond = 30
)

func Init(inf *infra.Infra) UserRepository {
	repo := &repository{inf.Db, inf.Redis}
	return repo
}

func (r *repository) Insert(user *model.InsertUserDb) (int64, error) {
	stmt, err := r.Db.Prepare(db.QueryUserInsert)
	if err != nil {
		log.Errorf("failed to prepare db statement: %v", err)
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Errorf("failed to close statement: %v", err)
		}
	}(stmt)
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		log.Errorf("failed to execute db statement: %v", err)
		return 0, err
	}
	log.Infof("successfully insert data to users")
	return res.LastInsertId()
}

func (r *repository) FindByEmail(email string) (*model.FindUserDb, error) {
	redisKey := redis2.GetKeyUserFindByEmail(email)
	cache, _ := r.findCache(redisKey)
	if cache != nil {
		return cache, nil
	}
	data, err := r.find(db.QueryUserFindByEmail, email)
	if err != nil {
		return nil, err
	}
	_ = r.setCache(redisKey, *data)
	return data, nil
}

func (r *repository) FindById(id int64) (*model.FindUserDb, error) {
	redisKey := redis2.GetKeyUserFindById(id)
	cache, _ := r.findCache(redisKey)
	if cache != nil {
		return cache, nil
	}
	data, err := r.find(db.QueryUserFindById, id)
	if err != nil {
		return nil, err
	}
	_ = r.setCache(redisKey, *data)
	return data, nil
}

func (r *repository) findCache(key string) (*model.FindUserDb, error) {
	user := &model.FindUserDb{}
	err := redis2.Get(r.Redis, key, &user)
	if err != nil {
		log.Warnf("redis cache not found")
		return nil, err
	}
	log.Infof("redis cache found with key %s", key)
	return user, nil
}

func (r *repository) setCache(key string, value model.FindUserDb) error {
	err := redis2.Set(r.Redis, key, value, userCacheExpirationSecond)
	if err != nil {
		log.Warnf("failed to set redis cache")
		return err
	}
	log.Infof("success set redis cache with key %s", key)
	return nil
}

func (r *repository) find(query string, arg interface{}) (*model.FindUserDb, error) {
	res := model.FindUserDb{}
	rows, err := r.Db.Query(query, arg)
	if err != nil {
		log.Errorf("failed query to database: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Errorf("failed to close rows: %v", err)
		}
	}(rows)
	exist := rows.Next()
	if !exist {
		log.Warnf("data not exists on database")
		return nil, errors.New("data not exists on database")
	}
	err = rows.Scan(&res.Id, &res.Name, &res.Email, &res.Password)
	if err != nil {
		log.Errorf("failed scan rows: %v", err)
		return nil, err
	}
	log.Infof("successfully find data from users")
	return &res, nil
}
