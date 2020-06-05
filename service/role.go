package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hulujia/form"
	"hulujia/model"
	"hulujia/repository"
	"hulujia/util/response"
	"hulujia/util/sqlcnd"
)

var RoleService = newRoleService()

func newRoleService() *roleService {
	return &roleService{}
}

type roleService struct {
}

// 获取角色列表
func (s *roleService) Lists(cnd *sqlcnd.SqlCnd) (list []model.Roles, paging *sqlcnd.Paging)  {
	return repository.RoleRepository.List(cnd)
}

// 获取角色
func (s *roleService) Get(field interface{}) *model.Roles  {
	return repository.RoleRepository.Get(field)
}

// 查看角色名称是否重复
func (s *roleService) ExistByRoleName(roleName string) bool {
	role := s.Get(roleName)
	if role != nil && role.Id > 0 {
		return true
	}
	return false
}

// 创建角色
func (s *roleService) Create(ctx *gin.Context, data form.RoleForm) *model.Roles  {
	return repository.RoleRepository.Create(map[string]interface{}{
		"role_name":  	data.RoleName,
		"perms": 		data.Perms,
		"desc":    		data.Desc,
	})
}

// 更新角色
func (s *roleService) Update(ctx *gin.Context, data form.RoleFormUpdate, id int) (bool, error)  {
	role := s.Get(id)
	if role == nil  {
		return false, errors.New(response.GetMsg(response.ERROR_NOT_FOUND,"角色"))
	}
	if s.ExistByRoleName(data.RoleName) && role.RoleName != data.RoleName {
		return false, errors.New(data.RoleName + "已经存在")
	}
	// 超级管理员的角色不能被修改
	if role.RoleName == repository.RoleName  && data.RoleName != repository.RoleName {
		return false, errors.New("超级管理员角色名称不能修改")
	}
	updateMap := map[string]interface{}{
		"role_name": data.RoleName,
		"desc": data.Desc,
		"perms": data.Perms,
	}
	// 更新角色
	res := repository.RoleRepository.Update(updateMap,id)
	repository.RoleRepository.UpdatePerms(role,data.Perms)
	return res, nil
}

// 删除管理员
func (s *roleService) Delete(id int) (bool bool, error error)  {
	role := s.Get(id)
	if role == nil {
		return true, errors.New(response.GetMsg(response.ERROR_USER_NOT_FOUND,"角色"))
	} else {
		if role.RoleName == repository.RoleName {
			return true, errors.New("超级管理员角色不能删除")
		}
	}
	return repository.RoleRepository.Delete(id),nil
}