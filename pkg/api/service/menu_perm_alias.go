package service

import (
	"github.com/CMLiang/zeus-admin/pkg/api/dao"
)

var menuPermAlias = dao.MenuPermAlias{}

//MenuPermAlias
type MenuPermAlias struct {
}

// GetByAlias - now just for `perm` middleware to automatically check alias record
//func (MenuPermAlias) GetByAlias (alias string) model.MenuPermAlias {
//	return menuPermAlias.GetByAlias(alias)
//}
