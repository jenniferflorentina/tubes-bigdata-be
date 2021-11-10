package services

import (
	"github.com/tubes-bigdata/domain"
	"github.com/tubes-bigdata/utils"
)

func CreateUser(user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func GetAll() ([]*domain.User, *utils.RestErr) {
	user, restErr := domain.GetAll()
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func FindUser(id string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Find(id)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func DeleteUser(id string) *utils.RestErr {
	restErr := domain.Delete(id)
	if restErr != nil {
		return restErr
	}
	return nil
}

func UpdateUser(id string, user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Update(id, user)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}
