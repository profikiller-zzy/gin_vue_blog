package image_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 上传文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 返回信息
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

	// 判断路径是否存在，如果不存在则创建
	if _, err := os.Stat(basePath); os.IsNotExist(err) { // 当前指定文件路径不存在
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}

	var upResList []FileUploadResponse = make([]FileUploadResponse, len(FileHeaderList))
	for index, FileHeader := range FileHeaderList {
		//global.Log.Info(fmt.Sprintf("上传图片成功，第%d张图片：Filename: %s, Size: %d bytes", index+1, FileHeader.Filename, FileHeader.Size))
		filePath := path.Join("uploads", FileHeader.Filename)
		// 判断文件大小是否大于指定最大文件大小
		if FileHeader.Size > (size << 20) {
			upResList[index] = FileUploadResponse{
				FileName:  FileHeader.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传图片大小大于设定大小，设定大小为 %d MB，当前图片大小为 %.3f MB", size, float64(FileHeader.Size)/(2<<20)),
			}
		} else {
			err = c.SaveUploadedFile(FileHeader, filePath)
			if err != nil { // 图片上传失败
				upResList[index] = FileUploadResponse{
					FileName:  FileHeader.Filename,
					IsSuccess: false,
					Msg:       fmt.Sprintf("上传图片保存到本地失败，错误信息:%s", err.Error()),
				}
			} else { // 上传成功
				upResList[index] = FileUploadResponse{
					FileName:  FileHeader.Filename,
					IsSuccess: true,
					Msg:       fmt.Sprintf("上传成功，当前图片大小为 %.3f MB", float64(FileHeader.Size)/(2<<20)),
				}
			}
		}
	}
	response.OKWithData(upResList, c)
}
