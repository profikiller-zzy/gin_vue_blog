package image_service

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/plugin/qiniu"
	"gin_vue_blog_AfterEnd/utils"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ImageWhiteList 图片上传白名单
var ImageWhiteList = []string{
	".jpg",
	".png",
	".apng",
	".jpeg",
	".tiff",
	".gif",
	".ico",
	".svg",
	".webp",
}

// ImageUploadService 处理图片文件上传的方法
func (ImageService) ImageUploadService(FileHeader *multipart.FileHeader, c *gin.Context) (req model.FileUploadResponse) {
	var basePath string = global.Config.SaveUpload.Path
	var size int64 = global.Config.SaveUpload.Size

	fileName := FileHeader.Filename
	ext := strings.ToLower(filepath.Ext(fileName))

	// 如果用户上传的文件不在白名单中，则直接判断下一个文件
	if !utils.IsInStringList(ext, ImageWhiteList) {
		req = GenerateFileUploadReq(fileName, false, fmt.Sprintf("上传文件类型错误，当前文件后缀为%s", ext))
		return req
	}

	filePath := path.Join(basePath, fileName)

	// 判断文件大小是否大于指定最大文件大小，大于则直接判断下一个文件
	if FileHeader.Size > (size << 20) {
		req = GenerateFileUploadReq(fileName, false, fmt.Sprintf("上传图片大小大于设定大小，设定大小为 %d MB，当前图片大小为 %.3f MB", size, float64(FileHeader.Size)/(2<<20)))
		return req
	}

	// 先调用Open函数打开`*multipart.FileHeader`对应的文件，用fileObj接收文件对应的`multipart.File`
	// 将文件中的内容读出，存入`[]byte`，便于调用函数将文件存入七牛云服务器
	fileObj, err := FileHeader.Open()
	fileObjContent, err := ioutil.ReadAll(fileObj)
	if err != nil {
		global.Log.Error(err.Error())
	}
	// 调用MD5对读取出的[]byte内容进行hash
	imageHash := utils.MD5(fileObjContent)

	// 去数据库中查询该数据是否存在，若存在则直接下一个文件，不存在则上传和入库
	fmt.Println(imageHash)
	var banner model.BannerModel
	err = global.Db.Take(&banner, "hash = ?", imageHash).Error
	if err == nil { // 数据库中存在这张图片
		req = GenerateFileUploadReq(banner.Path, false, "图片已存在")
		return req
	}

	// 上传图片文件至七牛云或本地
	// 是否启用七牛云服务器
	if global.Config.QiNiu.Enable { // 启用了七牛云服务器，则将图片上传至七牛云服务器
		// 这里的filePath是图片的URL
		filePath, err = qiniu.UploadFileToQiNiu(fileObjContent, "gvb", fileName)
	} else { // 未启用，将图片保存在本地
		// 判断路径是否存在，如果不存在则创建
		if _, err := os.Stat(basePath); os.IsNotExist(err) { // 当前指定文件路径不存在
			err = os.MkdirAll(basePath, fs.ModePerm)
			if err != nil {
				global.Log.Error(err.Error())
			}
		}
		err = c.SaveUploadedFile(FileHeader, filePath)
	}

	// 图片上传是否失败
	if err != nil { // 图片上传本地或云服务器失败
		global.Log.Error(err.Error())
		var msg string
		if global.Config.QiNiu.Enable {
			msg = fmt.Sprintf("上传图片保存到七牛云服务器失败，错误信息:%s", err.Error())
		} else {
			msg = fmt.Sprintf("上传图片保存到本地失败，错误信息:%s", err.Error())
		}
		req = GenerateFileUploadReq(fileName, false, msg)
		return req
	} else { // 图片已经上传成功
		// 将上传的图片记录存入数据库
		// 将上传的图片记录存入数据库
		var image = model.BannerModel{
			Path:             filePath,
			Hash:             imageHash,
			Name:             fileName,
			ImageStorageMode: ctype.Local,
		}
		if global.Config.QiNiu.Enable {
			image.ImageStorageMode = ctype.QiNiu
		}
		err = global.Db.Create(&image).Error
		if err != nil {
			global.Log.Error(fmt.Sprintf("图片文件写入数据库出错，报错信息:%s", err.Error()))
		}

		var msg string
		if global.Config.QiNiu.Enable {
			msg = fmt.Sprintf("图片上传七牛云服务器成功，当前图片大小为 %.3f MB", float64(FileHeader.Size)/(2<<20))
		} else {
			msg = fmt.Sprintf("图片上传至本地成功，当前图片大小为 %.3f MB", float64(FileHeader.Size)/(2<<20))
		}
		req = GenerateFileUploadReq(filePath, true, msg)
		err = fileObj.Close()
		if err != nil {
			global.Log.Warn(fmt.Sprintf("文件关闭失败，报错信息:%s", err.Error()))
		}
		return req
	}
}

func GenerateFileUploadReq(filePath string, isSuccess bool, msg string) (req model.FileUploadResponse) {
	return model.FileUploadResponse{
		FilePath:  filePath,
		IsSuccess: false,
		Msg:       msg,
	}
}
