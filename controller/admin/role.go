package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"hulujia/controller"
	"hulujia/form"
	"hulujia/service"
	"hulujia/util/response"
	"hulujia/util/sqlcnd"
)

type RoleController struct {
	controller.BaseController
}

// 列表
func (c *RoleController) List(ctx *gin.Context)  {
	page := com.StrTo(ctx.DefaultQuery("page","1")).MustInt()
	limit := com.StrTo(ctx.DefaultQuery("limit","15")).MustInt()

	conditions := sqlcnd.NewSqlCnd()

	if roleName := ctx.Query("role_name"); roleName != "" {
		conditions.Like("role_name",roleName)
	}
	list, paging := service.RoleService.Lists(conditions.Page(page, limit).Desc("id"))
	c.Success(ctx, &sqlcnd.PageResult{Results: list, Page: paging})
}

// 创建角色
func (c *RoleController) Create(ctx *gin.Context)  {
	var roleForm form.RoleForm
	if err := ctx.ShouldBind(&roleForm); err != nil {
		fmt.Println(err)
		c.Failed(ctx,response.ERROR)
		return
	}
	if c.Validate(ctx,roleForm) {
		if user := service.RoleService.Create(ctx,roleForm); user != nil {
			c.Message(ctx,"角色创建成功")
			return
		}
		c.FailedWithMsg(ctx,response.ERROR_CREATE_FAIL,"角色")
		return
	}
}

// 更新角色
func (c *RoleController) Update(ctx *gin.Context)  {
	id := com.StrTo(ctx.Param("id")).MustInt()
	var roleForm form.RoleFormUpdate
	if err := ctx.ShouldBind(&roleForm); err != nil {
		c.Failed(ctx,response.ERROR)
		return
	}
	if c.Validate(ctx,roleForm) {
		res, err := service.RoleService.Update(ctx,roleForm,id);
		if err != nil {
			c.ValidFailed(ctx,err.Error())
			return
		}
		if !res {
			c.FailedWithMsg(ctx,response.ERROR_UPDATE_FAIL,"角色")
			return
		}
		c.Message(ctx,"角色更新成功")
		return
	}
}

// 获取角色信息
func (c *RoleController) Show(ctx *gin.Context)  {
	general := form.GeneralGetId{ID:com.StrTo(ctx.Param("id")).MustInt()}
	if c.Validate(ctx, general) {
		user := service.RoleService.Get(general.ID)
		c.Success(ctx,user)
	}
}

// 删除管理员
func (c *RoleController) Destroy(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	deleted, err := service.RoleService.Delete(id)
	if err != nil {
		c.ValidFailed(ctx,err.Error())
		return
	}
	if !deleted {
		c.FailedWithMsg(ctx,response.ERROR_DELETE_FAIL,"角色")
		return
	}
	c.Message(ctx,"角色删除成功")
	return
}