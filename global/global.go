package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Log *zap.Logger
)
