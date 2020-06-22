package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"hulujia/controller"
	"hulujia/form"
	"hulujia/service"
	"hulujia/util/response"
	"hulujia/util/sqlcnd"
)

type UserController struct {
	controller.BaseController
}

// 根据token获取管理员信息
func (c *UserController) UserInfo(ctx *gin.Context) {
	user := service.UserService.GetCurrentUser(ctx)
	c.Success(ctx,user)
}

// 获取用户信息
func (c *UserController) Show(ctx *gin.Context) {
	general := form.GeneralGetId{ID:com.StrTo(ctx.Param("id")).MustInt()}
	if c.Validate(ctx, general) {
		user := service.UserService.GetById(general.ID)
		c.Success(ctx,user)
	}
}

// 获取管理员列表
func (c *UserController) List(ctx *gin.Context)  {
	page := com.StrTo(ctx.DefaultQuery("page","1")).MustInt()
	limit := com.StrTo(ctx.DefaultQuery("limit","15")).MustInt()

	conditions := sqlcnd.NewSqlCnd()

	if username := ctx.Query("name"); username != "" {
		conditions.Like("name",username)
	}
	if phone := ctx.Query("phone"); phone != "" {
		conditions.Eq("phone",phone)
	}
	list, paging := service.UserService.Lists(conditions.Page(page, limit).Desc("id"))
	c.Success(ctx, &sqlcnd.PageResult{Results: list, Page: paging})
}

// @Summary 创建管理员
// @Tags 后台
// @Description ""
// @Produce  json
// @Accept  multipart/form-data
// @consumes formData
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param phone formData string true "手机号"
// @Param roles  body array  false "角色"
// @Security Bearer
// @Success 200 {object} response.ResponseMessage	"success"
// @Failure 401 {object} response.ResponseMessage 	"Token解析失败"
// @Failure 422 {object} response.ResponseMessage 	"请求参数缺失或者不正确返回"
// @Router /api/admin/user [post]
func (c *UserController) Create(ctx *gin.Context) {
	var userForm form.UserForm
	if err := ctx.ShouldBind(&userForm); err != nil {
		c.Failed(ctx,response.INVALID_PARAMS)
		return
	}
	if c.Validate(ctx,userForm) {
		// 创建管理员
		if user := service.UserService.Create(ctx,userForm); user != nil {
			c.Message(ctx,"管理员创建成功")
			return
		}
		c.FailedWithMsg(ctx,response.ERROR_CREATE_FAIL,"管理员")
		return
	}
}

// 更新管理员
func (c *UserController) Update(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	var userForm form.UserFormUpdate
	if err := ctx.ShouldBind(&userForm); err != nil {
		c.Failed(ctx,response.INVALID_PARAMS)
		return
	}
	if c.Validate(ctx,userForm) {
		res, err := service.UserService.Update(userForm,id)
		if err != nil {
			c.ValidFailed(ctx,err.Error())
			return
		}
		if !res {
			c.FailedWithMsg(ctx,response.ERROR_UPDATE_FAIL,"管理员")
			return
		}
		c.Message(ctx,"管理员更新成功")
		return
	}
}

// 删除管理员
func (c *UserController) Destroy(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	deleted, err := service.UserService.Delete(id)
	if err != nil {
		c.ValidFailed(ctx,err.Error())
		return
	}
	if !deleted {
		c.FailedWithMsg(ctx,response.ERROR_DELETE_FAIL,"管理员")
		return
	}
	c.Message(ctx,"管理员删除成功")
	return
}
