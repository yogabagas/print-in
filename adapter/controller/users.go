package controller

import (
	"context"
	"github/yogabagas/print-in/domain/service"
	"github/yogabagas/print-in/service/users/usecase"
)

type UsersControllerImpl struct {
	usersSvc usecase.UsersService
}

type UsersController interface {
	CreateUsers(ctx context.Context, req service.CreateUsersReq) error
	Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error)
}

func NewUsersController(userSvc usecase.UsersService) UsersController {
	return &UsersControllerImpl{usersSvc: userSvc}
}

func (uc *UsersControllerImpl) CreateUsers(ctx context.Context, req service.CreateUsersReq) error {
	return uc.usersSvc.CreateUsers(ctx, req)
}

func (uc *UsersControllerImpl) Login(ctx context.Context, req service.LoginReq) (*service.LoginRes, error) {
	return uc.usersSvc.Login(ctx, req)
}
