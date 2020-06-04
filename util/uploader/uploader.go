package uploader

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var aliyun = NewAliyun()


// 阿里云oss
type aliyunOssUploader struct {
	once   sync.Once
	bucket *oss.Bucket
}

func NewAliyun() *aliyunOssUploader {
	return &aliyunOssUploader{
		once:   sync.Once{},
		bucket: nil,
	}
}