package service

import (
	"gin_example/common/errcode"
	"gin_example/common/response"
	"gin_example/common/validator"
	"gin_example/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateRoleRequest struct {
	Name string `json:"name" binding:"required,min=1,max=30" example:"漏洞管理员"` // 角色名
}

type UpdateRoleRequest struct {
	Name    *string `json:"name" binding:"omitempty,min=1,max=30" example:"漏洞管理员"` // 角色名
	Default *bool   `json:"default"  binding:"omitempty" example:"false"`          // 是否默认
}

type ListRoleRequest struct {
	Name    *string `form:"name" binding:"omitempty,min=1,max=30" example:"漏洞管理员"` // 角色名
	Default *bool   `form:"default"  binding:"omitempty" example:"false"`          // 是否默认
}

type RoleResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

func CreateRole(c *gin.Context) {
	param := CreateRoleRequest{}
	validator.BindAndValid(c, &param)
	model.Create(&model.Role{}, &param)
	response.NewResponse(c).ToResponse(nil)
	return
}

func GetRoles(c *gin.Context) {
	var roles []RoleResponse
	param := ListRoleRequest{}
	validator.BindAndValid(c, &param)
	count := model.List(c, &model.Role{}, &param, &roles)
	response.NewResponse(c).ToListResponse(roles, count)
	return
}

func GetRole(c *gin.Context) {
	role := RoleResponse{}
	model.Get(&model.Role{}, c.Param("id"), &role)
	response.NewResponse(c).ToResponse(role)
	return
}

func DeleteRole(c *gin.Context) {
	model.LogicDelete(&model.Role{}, c.Param("id"))
	response.NewResponse(c).ToResponse(nil)
	return
}

func UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errcode.ParamError)
	}
	param := UpdateRoleRequest{}
	validator.BindAndValid(c, &param)
	role := model.Role{BaseModel: &model.BaseModel{ID: uint32(id)}}
	model.Update(&role, &param)
	response.NewResponse(c).ToResponse(nil)
	return
}
