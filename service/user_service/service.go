package user_service

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/gommon/log"

	"go-vote/model"
	repo "go-vote/repository/user_repository"
	"go-vote/util/crypto"
)

type UserService interface {
	Register(req *model.RegisterUserReq) (*model.RegisterUserRes, error)
	GetProfile(id int64) (*model.GetProfileUserRes, error)
	Login(req *model.LoginUserReq) (*model.LoginUserRes, error)
}

type service struct {
	Repo repo.UserRepository
}

func Init(repo *repo.UserRepository) UserService {
	return &service{*repo}
}

func (s *service) Register(req *model.RegisterUserReq) (*model.RegisterUserRes, error) {
	res := &model.RegisterUserRes{}
	valid, err := req.Validate()
	if !valid {
		res.Response = model.Response{
			Status: http.StatusBadRequest,
		}
		return res, err
	}
	existing, _ := s.Repo.FindByEmail(req.Email)
	if existing != nil {
		log.Warnf("register user cancelled; user already exists")
		res.Response = model.Response{
			Status: http.StatusNotAcceptable,
		}
		return res, errors.New("user already exists")
	}
	pw, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		res.Response = model.Response{
			Status: http.StatusInternalServerError,
		}
		return res, err
	}
	insert := &model.InsertUserDb{
		Name:     sql.NullString{String: req.Name, Valid: req.Name != ""},
		Email:    req.Email,
		Password: pw,
	}
	id, err := s.Repo.Insert(insert)
	if err != nil {
		log.Errorf("failed to insert to repository: %v", err)
		res.Response = model.Response{
			Status: http.StatusInternalServerError,
		}
		return res, err
	}
	res = &model.RegisterUserRes{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
		Response: model.Response{
			Status: http.StatusCreated,
		},
	}
	log.Infof("successfully registered user with id %v", id)
	return res, nil
}

func (s *service) GetProfile(id int64) (*model.GetProfileUserRes, error) {
	res := &model.GetProfileUserRes{}
	if id < 1 {
		log.Warnf("user id should be greater than 0")
		res.Response = model.Response{Status: http.StatusBadRequest}
		return res, errors.New("user id should be greater than 0")
	}
	data, err := s.Repo.FindById(id)
	if err != nil {
		log.Errorf("error fetch user from repository: %v", err)
		res.Response = model.Response{Status: http.StatusInternalServerError}
		return res, err
	}
	res = &model.GetProfileUserRes{
		Response: model.Response{Status: http.StatusOK},
		Name:     data.Name.String,
		Email:    data.Email,
	}
	log.Infof("success fetch user from repository: %s", res.Email)
	return res, nil
}

func (s *service) Login(req *model.LoginUserReq) (*model.LoginUserRes, error) {
	//TODO implement me
	panic("implement me")
}
