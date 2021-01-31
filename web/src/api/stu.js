import service from '@/utils/request'

// @Tags Stu
// @Summary 创建Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "创建Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stu/createStu [post]
export const createStu = (data) => {
     return service({
         url: "/stu/createStu",
         method: 'post',
         data
     })
 }


// @Tags Stu
// @Summary 删除Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "删除Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stu/deleteStu [delete]
 export const deleteStu = (data) => {
     return service({
         url: "/stu/deleteStu",
         method: 'delete',
         data
     })
 }

// @Tags Stu
// @Summary 删除Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stu/deleteStu [delete]
 export const deleteStuByIds = (data) => {
     return service({
         url: "/stu/deleteStuByIds",
         method: 'delete',
         data
     })
 }

// @Tags Stu
// @Summary 更新Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "更新Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stu/updateStu [put]
 export const updateStu = (data) => {
     return service({
         url: "/stu/updateStu",
         method: 'put',
         data
     })
 }


// @Tags Stu
// @Summary 用id查询Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "用id查询Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stu/findStu [get]
 export const findStu = (params) => {
     return service({
         url: "/stu/findStu",
         method: 'get',
         params
     })
 }


// @Tags Stu
// @Summary 分页获取Stu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Stu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stu/getStuList [get]
 export const getStuList = (params) => {
     return service({
         url: "/stu/getStuList",
         method: 'get',
         params
     })
 }