package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

// UploadFileToQiNiu 上传文件至七牛云，返回文件的URL和报错
func UploadFileToQiNiu(data []byte, prefix string, fileName string) (filePath string, err error) {
	qiNiu := global.Config.QiNiu
	upToken := qiNiu.GetUpToken()
	if upToken == "" {
		return "", errors.New("请先配置七牛云的AccessKey和SecretKey！")
	}
	if float64(len(data))/(1<<20) > qiNiu.Size {
		return "", errors.New("上传文件大小大于设定大小，请检查文件是否小于5MB")
	}
	cfg := qiNiu.GetCfg()

	// 构建上传表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	// PutRet 为七牛标准的上传回复内容
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := fmt.Sprintf("%s/%s_%s", prefix, time.Now().Format("2006-01-02 15:04:05.000"), fileName)

	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s/%s", qiNiu.CDN, ret.Key), nil
}
