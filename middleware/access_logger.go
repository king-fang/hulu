package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"hulujia/config"
	"hulujia/util"
	"log"
	"path"
	"strings"
	"time"
)

type ResponseData struct {
	Code int
	Data interface{}
	Msg string
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerToFile() gin.HandlerFunc {

	go handleAccessChannel()

	return func(c *gin.Context) {

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now().UnixNano() / 1000000

		// 处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}
		var res ResponseData
		if responseBody != "" {
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Code
				responseMsg = res.Msg
				responseData = res.Data
			}
		}
		// 结束时间
		endTime := time.Now().UnixNano() / 1000000

		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}
		// 日志格式
		accessLogMap := make(map[string]interface{})
		claims := jwt.ExtractClaims(c)
		accessLogMap["request_time"]      = startTime
		accessLogMap["request_method"]    = c.Request.Method
		accessLogMap["request_uri"]       = c.Request.RequestURI
		accessLogMap["request_proto"]     = c.Request.Proto
		accessLogMap["request_ua"]        = c.Request.UserAgent()
		accessLogMap["request_referer"]   = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = c.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"]  = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)
		if claims != nil {
			accessLogMap["user_id"] = claims["id"]
		}
		accessLogJson := util.JsonEncode(accessLogMap)
		accessChannel <- accessLogJson
	}
}

func handleAccessChannel() {
	logFilePath := config.App.AccessLogPath
	logFileName := "hulujia-access.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	linkFileName := path.Join(logFilePath, strings.Split(logFileName, ".")[0])
	logf, err := rotatelogs.New(
		linkFileName + "-%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(15天)
		rotatelogs.WithMaxAge(15*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Fatalf("failed to create access: %s", err)
	}
	for accessLog := range accessChannel {
		_,_ = logf.Write([]byte(accessLog + "\n"))
	}
	return
}
