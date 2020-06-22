package database

import (
	"github.com/jinzhu/gorm"
	"hulujia/model"
)

func Roreign(db *gorm.DB)  {
	// 角色权限表
	db.Model(&model.RolePerms{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")

	// 用户角色中间表（多对多）
	db.Model(&model.UserRoles{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
}