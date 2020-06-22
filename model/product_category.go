package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	ID  					int		`gorm:"primary_key" json:"id"`
	CategoryName 			string	`gorm:"size:50;not null;comment:'分类名称';default:''" json:"category_name"`
	CategoryParentId		int		`gorm:"not null;comment:'分类父级ID';default:'0'" json:"category_parent_id"`
	CategoryLevel  			byte	`gorm:"type:tinyint;not null;comment:'分类等级';default:1" json:"category_level"`
	CategoryPath  			string	`gorm:"type:varchar(255);not null;comment:'该类目所有父类目-id-';default:''" json:"category_path"`
	CategoryImage 			string	`gorm:"type:varchar(150);not null;comment:'分类等级';default:''" json:"category_image"`
	CategorySort 			int		`gorm:"type:int(11);not null;comment:'分类排序';default:'0'" json:"category_sort"`
	CategoryIsShow 			int		`gorm:"type:tinyint(1);not null;comment:'0-不展示 1-展示';default:'0'" json:"category_is_show"`
	TimeModel
	DeletedTimeModel
}

// 创建之前钩子
func (p *ProductCategory) BeforeCreate(scope *gorm.Scope) (err error) {
	if p.CategoryParentId > 0 {
		// 找出父级数据
		parentCategory := ProductCategory{}
		scope.NewDB().First(&parentCategory,"id = ?", p.CategoryParentId)
		if parentCategory.ID > 0 { // 如果存在父级
			// 子级根据父级等级 + 1
			p.CategoryLevel = parentCategory.CategoryLevel + 1
			if parentCategory.CategoryPath == "" {
				p.CategoryPath  = fmt.Sprintf("%s-%d-",parentCategory.CategoryPath,parentCategory.ID)
			} else {
				p.CategoryPath  = fmt.Sprintf("%s%d-",parentCategory.CategoryPath,parentCategory.ID)
			}
		} else {
			p.CategoryParentId = 0
		}
	}
	return
}

func (p *ProductCategory) BeforeUpdate(scope *gorm.Scope) (err error) {
	if p.CategoryParentId > 0 {
		// 找出父级数据
		pCategory := ProductCategory{}
		scope.NewDB().First(&pCategory,"id = ?", p.CategoryParentId)
		// 如果父级id 与更新的数据父级id相同则不进行处理
		if (pCategory.ID > 0) {
			// 子级根据父级等级 + 1
			scope.SetColumn("CategoryLevel", pCategory.CategoryLevel + 1)
			if pCategory.CategoryPath == "" {
				scope.SetColumn("CategoryPath", fmt.Sprintf("%s-%d-",pCategory.CategoryPath,pCategory.ID))
			} else {
				scope.SetColumn("CategoryPath", fmt.Sprintf("%s%d-",pCategory.CategoryPath,pCategory.ID))
			}
			return
		}
	}
	scope.SetColumn("CategoryLevel", 1)
	scope.SetColumn("CategoryPath", "")
	return
}