package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"column:role_name;type:varchar(255);"`
	RoleNote string `gorm:"column:role_note;type:text;"`
}

func (r *Role) TableName() string {
	return "roles"
}
