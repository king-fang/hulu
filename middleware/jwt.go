package middleware

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	jwt3 "gopkg.in/dgrijalva/jwt-go.v3"
	"hulujia/config"
	"hulujia/model"
	"hulujia/service"
	"net/http"
	"time"
)

var (
	LoginStandard = 1
	//LoginOAuth    = 2
)

// 账号密码登录绑定验证数据
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func JwtAuth(LoginType int) *jwt.GinJWTMiddleware {
	jwtMiddleware := &jwt.GinJWTMiddleware{
		Realm: "Jwt",
		Key:           []byte(config.App.JwtSecret),			// jwt key
		Timeout:       time.Hour * 24 * 15,
		MaxRefresh:    time.Hour * 24 * 30,
		LoginResponse: LoginResponse,
		RefreshResponse: LoginResponse,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(model.UserClaims); ok {
				return jwt.MapClaims{
					"id": v.ID,
					"name": v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		// 该操作执行用户授权的动作
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return authenticatorByPassword(c)
		},
		// 该操作是在jwt验证之前调用，主要验证数据格式，也可以做其他验证
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(model.UserClaims); ok {
				return true
			}
			return false
		},
		// 该操作是在设置载荷的数据格式
		IdentityHandler: func(claims jwt3.MapClaims) interface{} {
			return model.UserClaims{
				Name: claims["name"].(string),
				ID:   int(claims["id"].(float64)),
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	return jwtMiddleware
}

// 登录成功后返回的信息
func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": map[string]interface{}{
			"token":  token,
			"expire": expire.Format("2006-01-02 15:04:05"),
		},
		"message": "success",
	})
}

// 账号密码登录
func authenticatorByPassword(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.Bind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	name := loginVals.Username
	password := loginVals.Password
	ok, err, user := service.UserService.VerifyAndReturnUserInfo(name, password)
	if ok {
		return model.UserClaims{
			ID:   user.ID,
			Name: user.Name,
		}, nil
	}
	return nil, err
}