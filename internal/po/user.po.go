package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column: uuid; type: varchar(36); not null; index: idx_uuid; unique;"`
	UserName string    `gorm:"column: user_name; type: varchar(255);"`
	IsActive bool      `gorm:"column: is_active; type: boolean; default: true;"`
	Roles    []Role    `gorm:"many2many: user_roles;"`
}

func (u *User) TableName() string {
	return "users"
}
