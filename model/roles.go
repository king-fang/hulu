package model

// 角色
type Roles struct {
	Id 		 int 			`gorm:"primary_key" json:"id"`
	RoleName string 		`gorm:"size:50;not null;comment:'角色名称'" json:"role_name"`
	Desc 	 string 		`gorm:"size:50;not null;comment:'角色描述';default:''" json:"desc"`
	Perms 	 *RolePerms		`gorm:"foreignkey:RoleId" json:"perms"`
	Users    []*User     	`gorm:"many2many:user_roles;" json:"users"`
	TimeModel
}

// 角色权限
type RolePerms struct {
	Id 		 int 		`gorm:"primary_key" json:"id"`
	RoleId 	 int 		`gorm:"type:int(11);not null;comment:'角色Id';default:0" json:"role_id"`
	Perms 	 string 	`gorm:"type:json;not null;comment:'角色权限';default:''" json:"perms"`
}