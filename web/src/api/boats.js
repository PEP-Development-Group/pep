import service from '@/utils/request'

// @Tags Boats
// @Summary 创建Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "创建Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /boats/createBoats [post]
export const createBoats = (data) => {
     return service({
         url: "/boats/createBoats",
         method: 'post',
         data
     })
 }


// @Tags Boats
// @Summary 删除Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "删除Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /boats/deleteBoats [delete]
 export const deleteBoats = (data) => {
     return service({
         url: "/boats/deleteBoats",
         method: 'delete',
         data
     })
 }

// @Tags Boats
// @Summary 删除Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /boats/deleteBoats [delete]
 export const deleteBoatsByIds = (data) => {
     return service({
         url: "/boats/deleteBoatsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Boats
// @Summary 更新Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "更新Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /boats/updateBoats [put]
 export const updateBoats = (data) => {
     return service({
         url: "/boats/updateBoats",
         method: 'put',
         data
     })
 }


// @Tags Boats
// @Summary 用id查询Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "用id查询Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /boats/findBoats [get]
 export const findBoats = (params) => {
     return service({
         url: "/boats/findBoats",
         method: 'get',
         params
     })
 }


// @Tags Boats
// @Summary 分页获取Boats列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Boats列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /boats/getBoatsList [get]
 export const getBoatsList = (params) => {
     return service({
         url: "/boats/getBoatsList",
         method: 'get',
         params
     })
 }