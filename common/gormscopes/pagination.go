package gormscopes

import (
	"gin_example/global"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Pagination(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil {
			offset = 0
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		if limit > global.PAGINATION_MAX_SIZE {
			limit = global.PAGINATION_MAX_SIZE
		}
		return db.Offset(offset).Limit(limit)
	}
}
