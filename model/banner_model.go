package model

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path             string                 `json:"path"`                                // 图片URL，如果存储在本地则为图片路径，存储在云服务器上则是图片链接
	Hash             string                 `json:"hash"`                                // 图片的Hash值，用以判断重复图片
	Name             string                 `gorm:"size:36" json:"name"`                 // 图片的名称
	ImageStorageMode ctype.ImageStorageMode `gorm:"default:1" json:"image_storage_mode"` // 图片的存储方式，可以存储在本地或七牛云服务器上
}

// BeforeDelete 钩子函数，删除BannerModel记录前自动调用
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageStorageMode == ctype.Local {
		// 本地图片，删除数据库存储记录，还需要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Warnln(fmt.Sprintf("本地删除图片失败，图片路径为：%s", b.Path))
			return err
		}
	}
	// 存储在云服务器上的图片则不用删除图片在云服务器上的存储
	return nil
}
