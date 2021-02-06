import service from '@/utils/request'

// @Tags Class
// @Summary 创建Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "创建Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cls/createClass [post]
export const createClass = (data) => {
     return service({
         url: "/Cls/createClass",
         method: 'post',
         data
     })
 }


// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cls/deleteClass [delete]
 export const deleteClass = (data) => {
     return service({
         url: "/Cls/deleteClass",
         method: 'delete',
         data
     })
 }

// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cls/deleteClass [delete]
 export const deleteClassByIds = (data) => {
     return service({
         url: "/Cls/deleteClassByIds",
         method: 'delete',
         data
     })
 }

// @Tags Class
// @Summary 更新Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "更新Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Cls/updateClass [put]
 export const updateClass = (data) => {
     return service({
         url: "/Cls/updateClass",
         method: 'put',
         data
     })
 }


// @Tags Class
// @Summary 用id查询Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "用id查询Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Cls/findClass [get]
 export const findClass = (params) => {
     return service({
         url: "/Cls/findClass",
         method: 'get',
         params
     })
 }


// @Tags Class
// @Summary 分页获取Class列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Class列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cls/getClassList [get]
 export const getClassList = (params) => {
     return service({
         url: "/Cls/getClassList",
         method: 'get',
         params
     })
 }