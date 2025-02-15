package service

import (
	"github.com/CMLiang/zeus-admin/pkg/api/dao"
	"github.com/CMLiang/zeus-admin/pkg/api/dto"
	"github.com/CMLiang/zeus-admin/pkg/api/log"
	"github.com/CMLiang/zeus-admin/pkg/api/model"
)

var userSecretDao = dao.UserSecretDao{}

// DomainService
type UserSecretService struct {
}

// InfoOfId - get role info by id
func (us UserSecretService) InfoOfId(dto dto.GeneralGetDto) model.UserSecret {
	return userSecretDao.Get(dto.Id)
}

// Create - create a new domain
func (us UserSecretService) Create(dto dto.UserSecretCreateDto) model.UserSecret {
	userSecret := model.UserSecret{
		User_id:      dto.User_id,
		Account_name: dto.Account_name,
		Secret:       dto.Secret,
	}
	c := userSecretDao.Create(&userSecret)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return userSecret
}
