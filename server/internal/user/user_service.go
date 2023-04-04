package user

import (
	"context"
	"server/util"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

//the parameters will be asked from service layer to the handler layer hence (create user req)

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	//TODO hash password

	u := &User{
		Username: req.Username,
		email:    req.email,
		Password: hashedPassword,
	}
	//once we find the user, we will call the repository
	//calling the repository we defined inside service, call createuser and pass pointer to user

	r, err := s.Repository.createUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		email:    r.email,
	}

	return res, nil
}
