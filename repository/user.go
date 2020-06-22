package repository

import (
	"hulujia/conn"
	"hulujia/model"
	"hulujia/util/sqlcnd"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

// 获取用户信息
func (d *userRepository) Get(field interface{}) *model.User {
	user := &model.User{}
	err := conn.DB().
		Where("id = ?", field).
		Or("name = ?", field).
		Or("phone = ?", field).
		Preload("Roles").
		Preload("Roles.Perms").
		First(user).Error;
	if err != nil {
		return nil
	}
	return user
}

// 获取用户列表
func (d *userRepository) List(cnd *sqlcnd.SqlCnd) (list []model.User, paging *sqlcnd.Paging) {
	cnd.With("Roles").Find(conn.DB(), &list)
	count := cnd.Count(conn.DB(), &model.User{})
	paging = &sqlcnd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// 创建管理员
func (d *userRepository) Create(data map[string]interface{}) *model.User {
	user := model.User{
		Name:      data["name"].(string),
		Password:  data["password"].(string),
		Phone:     data["phone"].(string),
	}
	if err := conn.DB().Create(&user).Error; err != nil {
		return nil
	}
	return &user
}

// 删除管理员
func (d *userRepository) Delete(id int) bool {
	if err := conn.DB().Unscoped().Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// 更新管理员
func (d *userRepository) Update(data map[string]interface{}, id int) bool  {
	if err := conn.DB().Model(&model.User{}).Where("id = ?",id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

// 创建并更新管理员角色
func (d *userRepository) UpdateOrCreateRole(roles []int, user *model.User) {
	if len(roles) > 0 {
		count := conn.DB().Model(user).Association("Roles").Count()
		var userRole []*model.UserRoles
		for _,r := range roles {
			userRole = append(userRole,&model.UserRoles{
				RolesId: r,
				UserId:user.ID,
			})
		}
		// 新增
		if count == 0 {
			conn.DB().Model(user).Association("Roles").Append(userRole)
		} else {
			// 修改
			conn.DB().Model(user).Association("Roles").Replace().Append(userRole)
		}
	} else {
		// 没有设置就删除
		conn.DB().Model(user).Association("Roles").Replace()
	}
}