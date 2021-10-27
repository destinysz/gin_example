package model

import (
	"fmt"
	"gin_example/common/errcode"
	"gin_example/common/gormscopes"
	"gin_example/global"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID         uint32    `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     bool      `gorm:"default:true " json:"status"`
}

func AddCustomCallbacks(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("SetTimeForCreateCallback", setTimeForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("setTimeForUpdateCallback", setTimeForUpdateCallback)

}

// Hooks是针对某个model，对所有model可以使用Callbacks
// https://gorm.io/zh_CN/docs/write_plugins.html
func setTimeForCreateCallback(db *gorm.DB) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(123, nowTime, reflect.TypeOf(nowTime))

	createTime := db.Statement.Schema.FieldsByName["CreateTime"]
	if _, isNone := createTime.ValueOf(db.Statement.ReflectValue); isNone {
		db.Statement.SetColumn("CreateTime", nowTime)
	}
	updateTime := db.Statement.Schema.FieldsByName["UpdateTime"]
	if _, isNone := updateTime.ValueOf(db.Statement.ReflectValue); isNone {
		db.Statement.SetColumn("UpdateTime", nowTime)
	}
}

func setTimeForUpdateCallback(db *gorm.DB) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := db.Statement.Schema.FieldsByName["UpdateTime"]
	if _, isNone := updateTime.ValueOf(db.Statement.ReflectValue); isNone {
		db.Statement.SetColumn("UpdateTime", nowTime)
	}
}

func Get(model, id, result interface{}) {
	// model result接受指针类型
	err := global.DB.Model(model).First(result, "id = ? AND status = true", id).Error
	if err == nil {
		return
	}
	switch err.Error() {
	case "record not found":
		panic(errcode.NotFound)
	default:
		global.Log.Error(err.Error())
		panic(errcode.ServerError)
	}
}

func Create(model, data interface{}) {
	// model data接受指针类型
	copier.Copy(model, data)
	err := global.DB.Create(model).Error
	if err != nil {
		global.Log.Error(err.Error())
		panic(errcode.ServerError)
	}
}

func Delete(model, id interface{}) {
	// model 接受指针类型
	err := global.DB.Model(model).Where("id = ?", id).Delete(model).Error
	if err != nil {
		global.Log.Error(err.Error())
		panic(errcode.ServerError)
	}
}

func LogicDelete(model, id interface{}) {
	// model 接受指针类型
	err := global.DB.Model(model).Where("id = ?", id).Update("status", false).Error
	if err != nil {
		global.Log.Error(err.Error())
		panic(errcode.ServerError)
	}
}

func Update(model, data interface{}) {
	// model data接受指针类型
	copier.Copy(model, data)
	err := global.DB.Updates(model).Error
	if err != nil {
		global.Log.Error(err.Error())
		panic(errcode.ServerError)
	}
}

func List(c *gin.Context, model, queue, result interface{}) (count int64) {
	// model queue result接受指针类型
	db := global.DB.Model(model).Where("status = true").Where(queue)
	db.Count(&count)
	db.Scopes(gormscopes.Pagination(c)).Find(result)
	return count
}
