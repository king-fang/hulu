package repository

import (
	"hulujia/model"
	"hulujia/util/sqlcnd"
)

var RoleRepository = newRoleRepository()

func newRoleRepository() *roleRepository {
	return &roleRepository{}
}

type roleRepository struct {
}

// 获取用户信息
func (d *roleRepository) Get(field interface{}) *model.Roles {
	role := &model.Roles{}
	err := db.Where("id = ?", field).Or("role_name = ?", field).Preload("Perms").First(role).Error;
	if err != nil {
		return nil
	}
	return role
}

// 获取角色列表
func (d *roleRepository) List(cnd *sqlcnd.SqlCnd) (list []model.Roles, paging *sqlcnd.Paging) {
	cnd.With("Perms").Find(db, &list)
	count := cnd.Count(db, &model.Roles{})
	paging = &sqlcnd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// 创建角色
func (d *roleRepository) Create(data map[string]interface{}) *model.Roles {
	role := model.Roles{
		RoleName: data["role_name"].(string),
		Desc: data["desc"].(string),
		Perms: &model.RolePerms{
			Perms: data["perms"].(string),
		},
	}
	if err := db.Create(&role).Error; err != nil {
		return nil
	}
	return &role
}

// 更新角色
func (d *roleRepository) Update(data map[string]interface{}, id int) bool  {
	if err := db.Model(&model.Roles{}).Where("id = ?",id).Updates(model.Roles{
		RoleName: data["role_name"].(string),
		Desc: data["desc"].(string),
	}).Error; err != nil {
		return false
	}
	return true
}

// 删除管理员
func (d *roleRepository) Delete(id int) bool {
	if err := db.Delete(&model.Roles{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// 更改角色权限关系
func (d *roleRepository) UpdatePerms( role *model.Roles, perm string) {
	db.Model(model.RolePerms{}).Where("role_id = ?", role.Id).Update(map[string]string{
		"perms": perm,
	})
}