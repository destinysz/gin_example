package model

import (
	"gin_example/global"
)

type Role struct {
	*BaseModel
	Name    *string `json:"name"`
	Default *bool   `json:"default"`
}

func (role Role) Create() error {
	return global.DB.Create(&role).Error
}
