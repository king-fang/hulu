package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"hulujia/util/strs"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// 是否是正式环境
func IsProd() bool {
	return viper.GetString("mode") == "release"
}

// index of
func IndexOf(userIds []int64, userId int64) int {
	if userIds == nil || len(userIds) == 0 {
		return -1
	}
	for i, v := range userIds {
		if v == userId {
			return i
		}
	}
	return -1
}

// 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
func IsValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// 验证是否是合法的邮箱
func IsValidateEmail(email string) (err error) {
	if len(email) == 0 {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

// 是否是合法的密码
func IsValidatePassword(password, rePassword string) error {
	if len(password) == 0 {
		return errors.New("请输入密码")
	}
	if strs.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}

// 是否是合法的URL
func IsValidateUrl(url string) error {
	if len(url) == 0 {
		return errors.New("URL格式错误")
	}
	indexOfHttp := strings.Index(url, "http://")
	indexOfHttps := strings.Index(url, "https://")
	if indexOfHttp == 0 || indexOfHttps == 0 {
		return nil
	}
	return errors.New("URL格式错误")
}

// 是否为手机号
func IsPhone(phone string) error {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
		return errors.New("手机号格式错误")
	}
	return nil
}

// 生成随机6位数字
func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

// 验证字符串是否为空
func IsEmpty(str string) bool  {
	if len(str) == 0 {
		return true
	}
	return false
}

// 验证是否为数字
func IsNumber(str string) bool  {
	if m, _ := regexp.MatchString("^[0-9]+$", str); !m {
		return false
	}
	return true
}

// 是否中文汉字
func IsCN(str string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", str); !m {
		return false
	}
	return true
}

// 是否英文
func IsEnglish(str string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		return false
	}
	return true
}

// json
func JsonEncode(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JsonDecode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}