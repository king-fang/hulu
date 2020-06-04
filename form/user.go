package form

type UserForm struct {
	Name 		string 	`form:"name" alias:"用户名" valid:"Required;Unique;" label:"user.name"`
	Password 	string	`form:"password" alias:"密码" valid:"Required;MinSize(6)"`
	Phone 		string 	`form:"phone" valid:"Required;Phone;Unique;" label:"user.phone"`
	Roles 		[]int 	`form:"roles[]"`
}

type UserFormUpdate struct {
	Name 		string 	`form:"name" alias:"用户名" valid:"Required;"`
	Password 	string	`form:"password"`
	Phone 		string 	`form:"phone" valid:"Required;Phone"`
	Roles 		[]int 	`form:"roles[]"`
}