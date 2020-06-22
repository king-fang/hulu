package form

type ProductCategoryForm struct {
	CateName 	string 	`form:"cate_name" alias:"分类名称" valid:"Required;"`
	ParentId 	int		`form:"parent_id"`
	Image 		string 	`form:"image"`
	Sort 		int 	`form:"sort"`
	IsShow 		int 	`form:"is_show"`
}

