package user_repository

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"go-vote/infra"

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

func Init(inf *infra.Infra) UserRepository {
	repo := &repository{inf.Db, inf.Redis}
	return repo
}

func (r *repository) Insert(user *model.InsertUserDb) (int64, error) {
	const query = "insert into users (name, email, password) values (?, ?, ?);"
	stmt, err := r.Db.Prepare(query)
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
	const query = "select id, name, email, password from users where email = ?;"
	return r.find(query, email)
}

func (r *repository) FindById(id int64) (*model.FindUserDb, error) {
	const query = "select id, name, email, password from users where password = ?;"
	return r.find(query, id)
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
