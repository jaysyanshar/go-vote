package user_service

import (
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"go-vote/config"
	"go-vote/model"
	repo "go-vote/repository/user_repository"
	"go-vote/util/crypto"
	"go-vote/util/jwt"
	"go-vote/util/response"
	"net/http"
)

type UserService interface {
	Register(req *model.RegisterUserReq) (*model.RegisterUserRes, error)
	GetProfile(id int64) (*model.GetProfileUserRes, error)
	Login(req *model.LoginUserReq) (*model.LoginUserRes, error)
	Refresh(req *model.RefreshUserReq) (*model.RefreshUserRes, error)
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
		res.Response = response.MakeResponse(http.StatusBadRequest)
		return res, err
	}
	existing, _ := s.Repo.FindByEmail(req.Email)
	if existing != nil {
		log.Warnf("register user cancelled; user already exists")
		res.Response = response.MakeResponse(http.StatusNotAcceptable)
		return res, errors.New("user already exists")
	}
	pw, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
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
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}
	res = &model.RegisterUserRes{
		Id:       id,
		Name:     req.Name,
		Email:    req.Email,
		Response: response.MakeResponse(http.StatusCreated),
	}
	log.Infof("successfully registered user with id %v", id)
	return res, nil
}

func (s *service) GetProfile(id int64) (*model.GetProfileUserRes, error) {
	res := &model.GetProfileUserRes{}
	if id < 1 {
		log.Warnf("user id should be greater than 0")
		res.Response = response.MakeResponse(http.StatusBadRequest)
		return res, errors.New("user id should be greater than 0")
	}
	data, err := s.Repo.FindById(id)
	if err != nil {
		log.Errorf("error fetch user from repository: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}
	res = &model.GetProfileUserRes{
		Response: response.MakeResponse(http.StatusOK),
		Name:     data.Name.String,
		Email:    data.Email,
	}
	log.Infof("success fetch user from repository: %s", res.Email)
	return res, nil
}

func (s *service) Login(req *model.LoginUserReq) (*model.LoginUserRes, error) {
	res := &model.LoginUserRes{}
	valid, err := req.Validate()
	if !valid {
		log.Warnf("invalid data on request: %v", err)
		res.Response = response.MakeResponse(http.StatusBadRequest)
		return res, err
	}

	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		log.Errorf("error get user data: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}

	validPassword := crypto.CheckPasswordHash(req.Password, user.Password)
	if !validPassword {
		log.Warnf("wrong password of user %v", user.Id)
		res.Response = response.MakeResponse(http.StatusNotAcceptable)
		return res, errors.New("invalid password")
	}

	token, err := jwt.CreateToken(user.ToUser(), req.IpAddress)
	if err != nil {
		log.Errorf("failed to create token: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}

	refresh, err := jwt.CreateRefreshToken(user.ToUser(), req.IpAddress)
	if err != nil {
		log.Errorf("failed to create refresh token: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}

	res.AccessToken = token
	res.RefreshToken = refresh
	log.Infof("login success with user id %v", user.Id)
	return res, nil
}

func (s *service) Refresh(req *model.RefreshUserReq) (*model.RefreshUserRes, error) {
	cfg := config.Get()
	res := &model.RefreshUserRes{}

	auth, err := jwt.ParseToken(req.RefreshToken, cfg.RefreshSecret)
	if err != nil {
		log.Errorf("failed to parse token: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}
	if auth.IpAddress != req.IpAddress {
		log.Warnf("ip address is different from existing")
		res.Response = response.MakeResponse(http.StatusUnauthorized)
		return res, errors.New("ip address is different from existing, please login again")
	}
	user := auth.ToUser()
	token, err := jwt.CreateToken(user, auth.IpAddress)
	if err != nil {
		log.Errorf("failed to create token: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}

	refresh, err := jwt.CreateRefreshToken(user, auth.IpAddress)
	if err != nil {
		log.Errorf("failed to create refresh token: %v", err)
		res.Response = response.MakeResponse(http.StatusInternalServerError)
		return res, err
	}

	res.AccessToken = token
	res.RefreshToken = refresh
	log.Infof("refresh success with user id %v", user.Id)
	return res, nil
}
