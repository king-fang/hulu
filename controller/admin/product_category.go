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

type ProductCategory struct {
	controller.BaseController
}

// 获取分类列表
func (c *ProductCategory) List(ctx *gin.Context)  {
	page := com.StrTo(ctx.DefaultQuery("page","1")).MustInt()
	limit := com.StrTo(ctx.DefaultQuery("limit","15")).MustInt()
	conditions := sqlcnd.NewSqlCnd()
	if categoryName := ctx.Query("category_name"); categoryName != "" {
		conditions.Like("category_name",categoryName)
	}
	list, paging := service.ProductCategoryService.Lists(conditions.Page(page, limit).Desc("id"))
	c.Success(ctx, &sqlcnd.PageResult{Results: list, Page: paging})
}

// 创建分类
func (c *ProductCategory) Create(ctx *gin.Context)  {
	var categoryForm form.ProductCategoryForm
	if err := ctx.ShouldBind(&categoryForm); err != nil {
		c.Failed(ctx,response.ERROR)
		return
	}
	if c.Validate(ctx,categoryForm) {
		if user := service.ProductCategoryService.Create(categoryForm); user != nil {
			c.Message(ctx,"分类创建成功")
			return
		}
		c.FailedWithMsg(ctx,response.ERROR_CREATE_FAIL,"分类")
		return
	}
}

// 更新分类
func (c *ProductCategory) Update(ctx *gin.Context)  {
	id := com.StrTo(ctx.Param("id")).MustInt()
	var categoryForm form.ProductCategoryForm
	if err := ctx.ShouldBind(&categoryForm); err != nil {
		c.Failed(ctx,response.ERROR)
		return
	}
	if c.Validate(ctx,categoryForm) {
		res, err := service.ProductCategoryService.Update(categoryForm,id)
		if err != nil {
			c.ValidFailed(ctx,err.Error())
			return
		}
		if res {
			c.FailedWithMsg(ctx,response.ERROR_UPDATE_FAIL,"分类")
			return
		}
		c.Message(ctx,"分类更新成功")
		return
	}
}

func (c *ProductCategory) Destroy(ctx *gin.Context)  {
	id := com.StrTo(ctx.Param("id")).MustInt()
	deleted, err := service.ProductCategoryService.Delete(id)
	if err != nil {
		c.ValidFailed(ctx,err.Error())
		return
	}
	if !deleted {
		c.FailedWithMsg(ctx,response.ERROR_DELETE_FAIL,"分类")
		return
	}
	c.Message(ctx,"分类删除成功")
	return
}