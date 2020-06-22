package database

import (
	"github.com/jinzhu/gorm"
	"hulujia/config"
	"hulujia/model"
	"hulujia/util"
)

// 创建一个超级管理员
func createUser(db *gorm.DB)  {
	user := &model.User{}
	err := db.First(&user,"id = ?", 1).Error
	if err != nil {
		db.Create(&model.User{
			Name:  		config.AdminName,
			Password:  	util.EncodePassword(config.AdminPassword),
			Phone:      config.AdminPhone,
		})
	}
}

// 创建一个超级角色
func createRole(db *gorm.DB)  {
	role := &model.Roles{}
	err := db.First(role,"role_name = ?", config.RoleName).Error
	if err != nil {
		db.Create(&model.Roles{
			RoleName:  	config.RoleName,
			Desc:      	"超级管理员",
		})
	}
}

// 管理员与角色关联
func createUserRoles(db *gorm.DB)  {
	user := &model.User{}
	err := db.First(&user,"id = ?", 1).Error
	if err == nil {
		count := db.Model(&user).Association("Roles").Count()
		if count == 0 {
			db.Model(user).Association("Roles").Append(model.UserRoles{RolesId: 1,UserId:user.ID})
		}
	}
}

// 迁移
func Seeder(db *gorm.DB)  {
	createUser(db)
	createRole(db)
	createUserRoles(db)
}
