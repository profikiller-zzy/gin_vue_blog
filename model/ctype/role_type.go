package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin      = 1
	PermissionUser       = 2
	PermissionVisitor    = 3
	PermissionDeniedUser = 4
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	switch r {
	case PermissionAdmin:
		return "PermissionAdmin"
	case PermissionUser:
		return "PermissionUser"
	case PermissionVisitor:
		return "PermissionVisitor"
	case PermissionDeniedUser:
		return "PermissionDeniedUser"
	default:
		return "OtherUser"
	}
}
