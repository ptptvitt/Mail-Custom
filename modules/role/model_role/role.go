package model_role

import (
	"task1/common"
)

type Role struct {
	common.SQLModel `json:",inline"`
	RoleCode        string       `json:"role_code" gorm:"column:role_code;primary_key"`
	RoleName        string       `json:"name_role" gorm:"column:name_role"`
	DepartmentId    int          `json:"department_id" gorm:"column:dept_id"`
	Permissions     []Permission `json:"permissions" gorm:"many2many:role_permission;joinForeignKey:role_code;joinReferences:permission_code"`
}

func (r Role) TableName() string {
	return "roles"
}
