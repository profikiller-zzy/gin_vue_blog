package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetMsgLabel 对参数进行校验，将产生error的字段中标签为`msg`的内容返回
func GetMsgLabel(err error, obj interface{}) string {
	// 使用时，应该传入结构体的指针
	objType := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个报错信息
			// 根据报错字段名，获取结构体的具体字段
			field := e.Field()
			if f, exits := objType.Elem().FieldByName(field); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
