package model

type User struct {
	ID        	int 		`gorm:"primary_key" json:"id"`
	Name 		string 		`gorm:"size:50;not null;comment:'用户名'" json:"name"`
	Password 	string 		`gorm:"size:150;not null;comment:'密码'" json:"-"`
	Phone 		string 		`gorm:"type:char(11);not null;comment:'手机号'" json:"phone"`
	Roles		[]*Roles 	`gorm:"many2many:user_roles" json:"roles"`
	TimeModel
	DeletedTimeModel
}


// 用户角色中间表
type UserRoles struct {
	UserId  	int `gorm:"type:int(11);unsigned"`
	RolesId 	int `gorm:"type:int(11);unsigned"`
}