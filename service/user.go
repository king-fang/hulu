package service

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"hulujia/config"
	"hulujia/form"
	"hulujia/model"
	"hulujia/repository"
	"hulujia/util"
	"hulujia/util/response"
	"hulujia/util/sqlcnd"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

// 根据id找出管理员信息
func (s *userService) GetById(id int) *model.User {
	return repository.UserRepository.Get(id)
}

// 根据name找出管理员信息
func (s *userService) GetByName(name string) *model.User {
	return repository.UserRepository.Get(name)
}

// 根据phone找出管理员信息
func (s *userService) GetByPhone(phone string) *model.User {
	return repository.UserRepository.Get(phone)
}

// 手机号是否存在
func (s *userService) ExistByPhone(phone string) bool {
	user := s.GetByPhone(phone)
	if user == nil {
		return false
	}
	return true
}

// 用户名是否存在
func (s *userService) ExistByName(name string) bool {
	user := s.GetByName(name)
	if user == nil {
		return false
	}
	return true
}

// 获取当前管理员登录用户
func (s *userService) GetCurrentUser(ctx *gin.Context) *model.User {
	claims := jwt.ExtractClaims(ctx)
	if claims == nil {
		return nil
	}
	user := repository.UserRepository.Get(claims["id"])
	if user == nil {
		return nil
	}
	return user
}

// 获取管理员列表
func (s *userService) Lists(cnd *sqlcnd.SqlCnd) (list []model.User, paging *sqlcnd.Paging)  {
	return repository.UserRepository.List(cnd)
}

// 创建管理员
func (s *userService) Create(ctx *gin.Context, data form.UserForm) *model.User  {
	user := repository.UserRepository.Create(map[string]interface{}{
		"name":    	data.Name,
		"password": util.EncodePassword(data.Password),
		"phone":    data.Phone,
	})
	// 创建管理员角色
	repository.UserRepository.UpdateOrCreateRole(data.Roles, user)
	return user
}

// 更新管理员
func (s *userService) Update(data form.UserFormUpdate, id int) (bool, error) {
	user := s.GetById(id)
	if user == nil {
		return false, errors.New(response.GetMsg(response.ERROR_USER_NOT_FOUND,"管理员"))
	}
	if s.ExistByName(data.Name) && user.Name != data.Name {
		return false, errors.New(response.GetMsg(response.ERROR_EXIST,"用户名"))
	}
	if s.ExistByPhone(data.Phone) && user.Phone != data.Phone {
		return false, errors.New(response.GetMsg(response.ERROR_EXIST,"手机号"))
	}
	if data.Password != "" && len(data.Password) < 6 {
		return false, errors.New("密码最少为6位")
	}
	updateData := make(map[string]interface{})
	updateData["name"] = data.Name
	if data.Password != "" {
		updateData["password"] = util.EncodePassword(data.Password)
	}
	updateData["phone"] = data.Phone
	res := repository.UserRepository.Update(updateData, id)
	repository.UserRepository.UpdateOrCreateRole(data.Roles, user)
	return res, nil
}

// 删除管理员
func (s *userService) Delete(id int) (bool bool, error error)  {
	user := s.GetById(id)
	if user == nil {
		return true, errors.New(response.GetMsg(response.ERROR_USER_NOT_FOUND,"管理员"))
	} else {
		if user.Name == config.AdminName && user.Phone == config.AdminPhone {
			return true, errors.New("超级管理员不能删除")
		}
	}
	return repository.UserRepository.Delete(id),nil
}



// 验证管理员账户并且返回用户信息
func (s *userService) VerifyAndReturnUserInfo(username, password string) (bool, error, model.User) {
	var userModel *model.User = nil
	userModel = repository.UserRepository.Get(username)

	if userModel == nil {
		return false, errors.New("账号或密码错误"), model.User{}
	}
	if userModel.ID < 1 {
		return false, errors.New("账号或密码错误"), model.User{}
	}
	if !util.ValidatePassword(userModel.Password, password) {
		//log.Error("password wrong: username=%s", userModel.Name)
		return false, errors.New("账号或密码错误"), model.User{}
	}
	return true, nil, *userModel
}
