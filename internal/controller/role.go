package controller

import (
	"gin_example/internal/service"

	"github.com/gin-gonic/gin"
)

type Role struct{}

func NewRole() Role {
	return Role{}
}

// @Summary 角色详情
// @Produce  json
// @Tags 角色
// @Router /roles/{id}/ [get]
// @Param id path int true "id"
// @Success 200 {object} service.RoleResponse
func (t Role) GetOne(c *gin.Context) {
	service.GetRole(c)
}

// @Summary 角色列表
// @Produce  json
// @Tags 角色
// @Param name query string false "角色名称" maxlength(30)
// @Param offset query int false "分页起始位置" default(0)
// @Param limit query int false "每页数量 最大100" default(10) maximum(100)
// @Router /roles/ [get]
// @Success 200 {object} service.RoleResponse
func (t Role) GetList(c *gin.Context) {
	service.GetRoles(c)
}

// @Summary 添加角色
// @Produce  json
// @Tags 角色
// @Param body body service.CreateRoleRequest true "body"
// @Router /roles/ [post]
// @Success 200 {string} json "{}"
func (t Role) Create(c *gin.Context) {
	service.CreateRole(c)
}

// @Summary 修改角色
// @Produce  json
// @Tags 角色
// @Param id path int true "id"
// @Param body body service.UpdateRoleRequest true "body"
// @Router /roles/{id}/ [put]
// @Success 200 {string} json "{}"
func (t Role) Update(c *gin.Context) {
	service.UpdateRole(c)
}

// @Summary 删除角色
// @Produce  json
// @Tags 角色
// @Router /roles/{id}/ [delete]
// @Param id path int true "id"
// @Success 200 {string} json "{}"
func (t Role) Delete(c *gin.Context) {
	service.DeleteRole(c)
}
func (t Role) AddUser(c *gin.Context)          {}
func (t Role) DeleteUser(c *gin.Context)       {}
func (t Role) GetUsers(c *gin.Context)         {}
func (t Role) AddPermission(c *gin.Context)    {}
func (t Role) DeletePermission(c *gin.Context) {}
func (t Role) Getpermissions(c *gin.Context)   {}
