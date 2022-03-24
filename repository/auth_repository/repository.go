package auth_repository

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"go-vote/infra"
	"go-vote/model"
)

type AuthRepository interface {
	Insert(auth *model.InsertAuthDb) (int64, error)
	FindById(id int64) (*model.FindAuthDb, error)
	UpdateExpired(auth *model.UpdateAuthDb) error
	RevokeById(id int64) error
}

type repository struct {
	Db    *sql.DB
	Redis *redis.Client
}

func Init(inf *infra.Infra) AuthRepository {
	repo := &repository{inf.Db, inf.Redis}
	return repo
}

func (r *repository) Insert(auth *model.InsertAuthDb) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindById(id int64) (*model.FindAuthDb, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateExpired(auth *model.UpdateAuthDb) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) RevokeById(id int64) error {
	//TODO implement me
	panic("implement me")
}
