package service

import (
	"errors"
	"hulujia/form"
	"hulujia/model"
	"hulujia/repository"
	"hulujia/util/sqlcnd"
)

var ProductCategoryService = newCategoryService()

func newCategoryService() *categoryService {
	return &categoryService{}
}

type categoryService struct {
}

// 获取分类信息
func (s *categoryService) Get(id int) *model.ProductCategory  {
	return repository.ProductCategoryRepository.Get(id)
}

// 获取分类列表
func (s *categoryService) Lists(cnd *sqlcnd.SqlCnd) (list []model.ProductCategory, paging *sqlcnd.Paging)  {
	return repository.ProductCategoryRepository.List(cnd)
}

// 创建分类
func (s *categoryService) Create(data form.ProductCategoryForm) *model.ProductCategory  {
	category := repository.ProductCategoryRepository.Create(map[string]interface{}{
		"cate_name":    	data.CateName,
		"parent_id":		data.ParentId,
		"image":    		data.Image,
		"sort":    			data.Sort,
		"is_show":    		data.IsShow,
	})
	return category
}

// 更新分类
func (s *categoryService) Update(data form.ProductCategoryForm, id int) (bool, error)  {
	// 如果该分类存在子级，
	if len(repository.ProductCategoryRepository.GetChildren(id)) > 0 && data.ParentId > 0{
		return false, errors.New("该分类下存在字级，不允许设置父级")
	}
	res := repository.ProductCategoryRepository.Update(map[string]interface{}{
		"category_name":    		data.CateName,
		"category_parent_id":		data.ParentId,
		"category_image":    		data.Image,
		"category_sort":    		data.Sort,
		"category_is_show":    		data.IsShow,
	},id)
	return res, nil
}

func (s *categoryService) Delete(id int) (bool, error)  {
	// 查询是否有子级分类
	if len(repository.ProductCategoryRepository.GetChildren(id)) > 0 {
		return false, errors.New("该分类存在子级分类，请先删除子级分类")
	}
	return repository.ProductCategoryRepository.Delete(id),nil
}

