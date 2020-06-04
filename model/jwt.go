package model

// Password 用户 jwt (账号密码专用)
type UserClaims struct {
	ID   int  `json:"id"`
	Name string `json:"name"`
}


