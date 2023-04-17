package image_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
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

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 上传文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 返回信息
}

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

// ImageUploadingView 上传图片并将图片保存在uploads文件夹中
func (ImageApi) ImageUploadingView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.LogFail(err, c)
		return
	}
	var FileHeaderList []*multipart.FileHeader = form.File["image"]
	if len(FileHeaderList) == 0 {
		response.FailWithMessage("没有指定任何文件或者文件不存在", c)
		return
	}
	var basePath string = global.Config.SaveUpload.Path
	var size int64 = global.Config.SaveUpload.Size

	var upResList []FileUploadResponse = make([]FileUploadResponse, len(FileHeaderList))
	for index, FileHeader := range FileHeaderList {
		fileName := FileHeader.Filename
		ext := strings.ToLower(filepath.Ext(fileName))

		// 如果用户上传的文件不在白名单中，则直接判断下一个文件
		if !utils.IsInStringList(ext, ImageWhiteList) {
			upResList[index] = FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传文件类型错误，当前文件后缀为%s", ext),
			}
			continue
		}

		filePath := path.Join(basePath, fileName)

		// 判断文件大小是否大于指定最大文件大小，大于则直接判断下一个文件
		if FileHeader.Size > (size << 20) {
			upResList[index] = FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传图片大小大于设定大小，设定大小为 %d MB，当前图片大小为 %.3f MB", size, float64(FileHeader.Size)/(2<<20)),
			}
			continue
		}

		// 先调用Open函数打开`*multipart.FileHeader`对应的文件，用fileObj接收文件对应的`multipart.File`
		// 将文件中的内容读出，存入`[]byte`，便于调用函数将文件存入七牛云服务器
		fileObj, err := FileHeader.Open()
		fileObjContent, err := ioutil.ReadAll(fileObj)
		if err != nil {
			global.Log.Error(err.Error())
		}

		imageHash := utils.MD5(fileObjContent)

		// 去数据库中查询该数据是否存在，若存在则跳过，不存在则上传加入库
		fmt.Println(imageHash)
		var banner model.BannerModel
		err = global.Db.Take(&banner, "hash = ?", imageHash).Error
		if err == nil { // 数据库中存在这张图片
			upResList[index] = FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       "图片已存在",
			}
			continue
		}

		// 是否启用七牛云服务器
		if global.Config.QiNiu.Enable { // 启用了七牛云服务器，则将图片上传至七牛云服务器
			filePath, err = qiniu.UploadFileToQiNiu(fileObjContent, "gvb", fileName)
			if err != nil {
				global.Log.Error(err.Error())
				continue
			}
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
		if err != nil { // 图片上传失败
			var msg string
			if global.Config.QiNiu.Enable {
				msg = fmt.Sprintf("上传图片保存到七牛云服务器失败，错误信息:%s", err.Error())
			} else {
				msg = fmt.Sprintf("上传图片保存到本地失败，错误信息:%s", err.Error())
			}
			upResList[index] = FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       msg,
			}
		} else { // 图片上传成功
			// 将上传的文件存入数据库
			// 调用MD5对读取出的[]byte内容进行hash

			// 不存在，则入库
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

			upResList[index] = FileUploadResponse{
				FileName:  fileName,
				IsSuccess: true,
				Msg:       fmt.Sprintf("上传成功，当前图片大小为 %.3f MB", float64(FileHeader.Size)/(2<<20)),
			}
			err = fileObj.Close()
			if err != nil {
				global.Log.Warn(fmt.Sprintf("文件关闭失败，报错信息:%s", err.Error()))
			}
		}
	}
	response.OKWithData(upResList, c)
}
