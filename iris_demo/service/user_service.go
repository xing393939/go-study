package service

import (
	"iris_demo/models"
	"iris_demo/repo"
)

type UserService interface {
	GetUserList() *models.Result
	PostSaveUser(user models.User) (result models.Result)
	GetUserById(id uint) (result models.Result)
	DelUser(id uint) (result models.Result)
}

type userService struct{
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc userService) GetUserList() *models.Result {
	books := svc.repo.GetUserList()
	result := new(models.Result)
	result.Data = books
	result.Code = 200
	result.Msg = "SUCCESS"
	return result
}

func (svc userService) PostSaveUser(user models.User) (result models.Result) {
	err := svc.repo.SaveUser(user)
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "SUCCESS"
		user := svc.repo.GetUserByName(user.Name)
		result.Data = user
	}
	return
}

func (svc userService) GetUserById(id uint) (result models.Result) {
	user, err := svc.repo.GetUserById(id)
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
	} else {
		result.Data = user
		result.Code = 200
		result.Msg = "SUCCESS"
	}
	return result
}

func (svc userService) DelUser(id uint) (result models.Result) {
	err := svc.repo.DeleteUser(id)
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "SUCCESS"
		list := svc.repo.GetUserList()
		result.Data = list
	}
	return
}
