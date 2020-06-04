package form

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"hulujia/repository"
	"strings"
)

type Tables struct {
	Table string
	Field string
}

func SetUp()  {
	var MessageTmpls = map[string]string{
		"Required":     "不能为空",
		"Min":          "最小为 %d",
		"Max":          "最大为 %d",
		"Range":        "范围在 %d 至 %d",
		"MinSize":      "最小长度为 %d",
		"MaxSize":      "最大长度为 %d",
		"Length":       "长度必须是 %d",
		"Alpha":        "必须是有效的字母字符",
		"Numeric":      "必须是有效的数字字符",
		"AlphaNumeric": "必须是有效的字母或数字字符",
		"Match":        "必须匹配格式 %s",
		"NoMatch":      "必须不匹配格式 %s",
		"AlphaDash":    "必须是有效的字母或数字或破折号(-_)字符",
		"Email":        "必须是有效的邮件地址",
		"IP":           "必须是有效的IP地址",
		"Base64":       "必须是有效的base64字符",
		"Mobile":       "必须是有效手机号码",
		"Tel":          "必须是有效电话号码",
		"Phone":        "必须是有效的电话号码或者手机号码",
		"ZipCode":      "必须是有效的邮政编码",
	}

	validation.SetDefaultMessage(MessageTmpls)

	//增加默认的自定义验证方法
	_ = validation.AddCustomFunc("Unique", Unique)
}

func Unique(v *validation.Validation, obj interface{}, key string) {
	strArrayNew := strings.Split(key,".")[2:]
	table := strArrayNew[0]		// 表
	field := strArrayNew[1]		// 字段
	query := fmt.Sprintf("%s = ?", field)
	rows, _ := repository.DB().Table(table).Where(query,obj).Rows()
	if (rows.Next()) {
		v.AddError(key, fmt.Sprintf("%s 已经存在",obj))
	}
}

