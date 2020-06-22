package database

import (
	"github.com/jinzhu/gorm"
	"hulujia/model"
)

var models = []interface{}{
	&model.User{},            // 管理员
	&model.Roles{},           // 角色
	&model.RolePerms{},       // 角色权限
	&model.ProductCategory{}, // 产品分类
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(models...).Error
}