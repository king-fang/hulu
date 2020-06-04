package form

type RoleForm struct {
	RoleName 	string 	`form:"role_name" alias:"角色名称" valid:"Required;Unique;" label:"roles.role_name" json:"role_name"`
	Perms 		string	`form:"perms" alias:"角色权限" valid:"Required"`
	Desc 		string 	`form:"desc"`
}

type RoleFormUpdate struct {
	RoleName 	string 	`form:"role_name" alias:"角色名称" valid:"Required;" json:"role_name"`
	Perms 		string	`form:"perms" alias:"角色权限" valid:"Required"`
	Desc 		string 	`form:"desc" json:"desc"`
}

