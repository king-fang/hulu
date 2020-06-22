package config

import (
	"fmt"
	"github.com/unknwon/com"
	"hulujia/util/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	AdminName = "admin"
	AdminPhone = "18507193432"
	AdminPassword = "password"
	RoleName = "super_admin"
	// 配置文件
	conf = "./app.yaml"
	// 超时时间
	app_read_time_out  = 120
	app_write_time_out = 120
)


type application struct {
	Name 				string			// 项目名称
	Port 				int				// 端口
	Url					string			// Url
	RunMode 			string			// 当前运行环境
	AccessLogPath 		string			// 请求日志
	JwtSecret 			string			// jwt盐
	ErrorNotifyEmail 	string			// 接受异常的邮箱
	AppReadTimeout  	time.Duration	// 超时时间
	AppWriteTimeout 	time.Duration   // 超时时间
}

type database struct {
	Type string				// 数据库类型 mysql
	User string				// 用户名
	Password string			// 密码
	Host string				// 链接地址
	Name string				// 数据库名
	Ssl bool				// 是否ssl
	Min int					// 最小链接数
	Max int					// 最大链接数
	CharSet	string			// 字符集
}



var App = &application{}
var Database = &database{}

func SetupConfig()  {
	// 读取配置
	viper.SetConfigFile(conf)
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read conf file fail: %s", err.Error()))
	}
	// 解析配置
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse conf file fail: %s", err.Error()))
	}

	// 1. 基本配置
	App.Name 				= viper.GetString("base.app_name")
	App.Port 				= com.StrTo(viper.GetString("base.port")).MustInt()
	App.Url 				= viper.GetString("base.url")
	App.RunMode 			= viper.GetString("mode")
	App.AccessLogPath 		= viper.GetString("base.access_log_path")
	App.JwtSecret 			= viper.GetString("base.jwt_key")
	App.ErrorNotifyEmail 	= viper.GetString("base.error_notify_email")
	App.AppReadTimeout 		= time.Duration(app_read_time_out) * time.Second
	App.AppWriteTimeout 	= time.Duration(app_write_time_out) * time.Second

	// 2.数据库配置
	Database.Type 		= viper.GetString("database.driver")
	Database.Host 		= viper.GetString("database.mysql.host")
	Database.User 		= viper.GetString("database.mysql.user")
	Database.Password 	= viper.GetString("database.mysql.password")
	Database.Name 		= viper.GetString("database.mysql.name")
	Database.CharSet 	= viper.GetString("database.mysql.charset")
	Database.Max 		= viper.GetInt("database.mysql.pool.max")
	Database.Min 		= viper.GetInt("database.mysql.pool.min")
	Database.Ssl 		= viper.GetBool("database.mysql.ssl")

	log.Info("Load conf file %s successfully", conf)
}

