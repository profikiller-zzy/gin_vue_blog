package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin      Role = 1 // 超级管理员
	PermissionUser       Role = 2 // 普通用户
	PermissionVisitor    Role = 3 // 游客
	PermissionDeniedUser Role = 4 // 被禁言的用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDeniedUser:
		str = "被禁用的用户"
	default:
		str = "其他"
	}
	return str
}

//func (r *Role) Scan(value interface{}) error {
//	v, _ := value.([]byte)
//
//	if strings.EqualFold(string(v), "管理员") {
//		*r = PermissionAdmin
//	} else if strings.EqualFold(string(v), "用户") {
//		*r = PermissionUser
//	} else if strings.EqualFold(string(v), "游客") {
//		*r = PermissionVisitor
//	} else if strings.EqualFold(string(v), "被禁用的用户") {
//		*r = PermissionDeniedUser
//	} else {
//		*r = PermissionUser
//	}
//
//	return nil
//}
//
//func (r Role) Value() (driver.Value, error) {
//	var str string
//	switch r {
//	case PermissionAdmin:
//		str = "管理员"
//	case PermissionUser:
//		str = "用户"
//	case PermissionVisitor:
//		str = "游客"
//	case PermissionDeniedUser:
//		str = "被禁用的用户"
//	default:
//		str = "其他"
//	}
//	return str, nil
//}
