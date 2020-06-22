package repository

import (
	"hulujia/conn"
	"hulujia/model"
	"hulujia/util/sqlcnd"
)

var ProductCategoryRepository = newCategoryRepository()

func newCategoryRepository() *cateRepository {
	return &cateRepository{}
}

type cateRepository struct {
}

// 获取分类
func (d *cateRepository) Get(field interface{}) *model.ProductCategory {
	category := &model.ProductCategory{}
	err := conn.DB().Where("id = ?", field).Or("category_name = ?", field).First(category).Error;
	if err != nil {
		return nil
	}
	return category
}

// 获取子级
func (d *cateRepository) GetChildren(id int) []model.ProductCategory {
	category := []model.ProductCategory{}
	conn.DB().Where("category_parent_id = ?", id).Find(&category)
	return category
}

// 获取所有分类
func (d *cateRepository) List(cnd *sqlcnd.SqlCnd) (list []model.ProductCategory, paging *sqlcnd.Paging) {
	cnd.Find(conn.DB(), &list)
	count := cnd.Count(conn.DB(), &model.User{})
	paging = &sqlcnd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// 增
func (d *cateRepository) Create(data map[string]interface{}) *model.ProductCategory {
	category := model.ProductCategory{
		CategoryName:     data["cate_name"].(string),
		CategoryParentId: data["parent_id"].(int),
		CategoryImage:    data["image"].(string),
		CategorySort:     data["sort"].(int),
		CategoryIsShow:   data["is_show"].(int),
	}
	if err := conn.DB().Create(&category).Error; err != nil {
		return nil
	}
	return &category
}

// 更
func (d *cateRepository) Update(data map[string]interface{}, id int) bool  {
	if err := conn.DB().Model(&model.ProductCategory{}).Where("id = ?",id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

// 删
func (d *cateRepository) Delete(id int) bool {
	if err := conn.DB().Delete(&model.ProductCategory{}, "id = ?", id).Error; err != nil {
		return false
	}
	return  true
}



