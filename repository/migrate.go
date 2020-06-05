package repository

import (
	"hulujia/model"
	"hulujia/util"
)

// 超级管理员设置
var (
	AdminName = "admin"
	AdminPhone = "18507193432"
	AdminPassword = "password"
	RoleName = "super_admin"
)

// 创建一个超级管理员
func CreateUser()  {
	admin := UserRepository.Get(1)
	if admin == nil {
		DB().Create(&model.User{
			Name:  		AdminName,
			Password:  	util.EncodePassword(AdminPassword),
			Phone:      AdminPhone,
		})
	}
}

// 创建一个超级角色
func CreateRole()  {
	role := RoleRepository.Get(RoleName)
	if role == nil {
		DB().Create(&model.Roles{
			RoleName:  	RoleName,
			Desc:      	"超级管理员",
		})
	}
}

// 管理员与角色关联
func CreateUserRoles()  {
	admin := UserRepository.Get(1)
	count := DB().Model(&admin).Association("Roles").Count()
	if count == 0 {
		DB().Model(admin).Association("Roles").Append(model.UserRoles{RolesId:1,UserId:admin.ID})
	}
}
