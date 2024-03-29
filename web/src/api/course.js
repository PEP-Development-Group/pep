import service from '@/utils/request'
//课程管理相关的API集中在此处

// @Tags Class
// @Summary 选课
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body 参数就两个，发群里了 true "选课"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"选课成功"}"
// @Router /class/selectClass [post]
export const selectClass = (data) => {
    return service({
        url: "/class/selectClass",
        method: 'post',
        data
    })
}

// @Tags Class
// @Summary 退课
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body 参数就两个，发群里了 true "退课"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"退课成功"}"
// @Router /class/deleteSelect [delete]
export const deleteSelect = (data) => {
    return service({
        url: "/class/deleteSelect",
        method: 'delete',
        data
    })
}

// @Tags Class
// @Summary 创建Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "创建Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/createClass [post]
export const createClass = (data) => {
    return service({
        url: "/class/createClass",
        method: 'post',
        data: data
    })
}


// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
export const deleteClass = (data) => {
    return service({
        url: "/class/deleteClass",
        method: 'delete',
        data: data
    })
}

// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
export const deleteClassByIds = (data) => {
    return service({
        url: "/class/deleteClassByIds",
        method: 'delete',
        data: data
    })
}

// @Tags Class
// @Summary 更新Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "更新Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /class/updateClass [put]
export const updateClass = (data) => {
    return service({
        url: "/class/updateClass",
        method: 'put',
        data: data
    })
}


// @Tags Class
// @Summary 用id查询Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "用id查询Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /class/findClass [get]
export const findClass = (params) => {
    return service({
        url: "/class/findClass",
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
// @Router /class/getClassList [get]
export const getClass = (params) => {
    return service({
        url: "/class/getClass",
        method: 'get',
        params
    })
}

export const GetClassListWithPerson = (data) => {
    return service({
        url: "/class/getClassList",
        method: 'get',
        data: data,
        donNotShowLoading: true
    })
}

export const SelectClass = (data) => {
    return service({
        url: "/class/selectClass",
        method: 'post',
        data: data,
        donNotShowLoading: false
    })
}

export const DeleteSelect = (data) => {
    return service({
        url: "/class/deleteSelect",
        method: 'delete',
        data: data,
        donNotShowLoading: false
    })
}

export const getUserCreditInfo = (params) => {
    return service({
        url: "/class/getUserCreditInfo",
        method: 'get',
        donNotShowLoading: true,
        params
    })
}

export const GetPersonalClasses = (params) => {
    return service({
        url: "/class/getPersonalClasses",
        method: 'get',
        donNotShowLoading: true,
        params
    })
}

export const getTeacherClassList = (params) => {
    return service({
        url: "/class/getTeacherClassList",
        method: 'get',
        donNotShowLoading: true,
        params
    })
}

export const getTeacherAClassStuList = (params) => {
    return service({
        url: "/class/getTeacherAClassStuList",
        method: 'get',
        donNotShowLoading: true,
        params
    })
}

export const setStuGrade = (data) => {
    return service({
        url: "/class/setStuGrade",
        method: 'put',
        data: data,
        donNotShowLoading: true
    })
}


export const addUserCancelNums = (data) => {
    return service({
        url: "/class/addUserCancelNums",
        method: 'put',
        data: data,
        donNotShowLoading: true
    })
}

// @Tags ClassList
// @Summary 创建ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "创建ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classlist/createClassList [post]
export const createClassList = (data) => {
    return service({
        url: "/classlist/createClassList",
        method: 'post',
        data
    })
}


// @Tags ClassList
// @Summary 删除ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "删除ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classlist/deleteClassList [delete]
export const deleteClassList = (data) => {
    return service({
        url: "/classlist/deleteClassList",
        method: 'delete',
        data
    })
}

// @Tags ClassList
// @Summary 删除ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classlist/deleteClassList [delete]
export const deleteClassListByIds = (data) => {
    return service({
        url: "/classlist/deleteClassListByIds",
        method: 'delete',
        data
    })
}

// @Tags ClassList
// @Summary 更新ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "更新ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /classlist/updateClassList [put]
export const updateClassList = (data) => {
    return service({
        url: "/classlist/updateClassList",
        method: 'put',
        data
    })
}


// @Tags ClassList
// @Summary 用id查询ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "用id查询ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /classlist/findClassList [get]
export const findClassList = (params) => {
    return service({
        url: "/classlist/findClassList",
        method: 'get',
        params
    })
}


// @Tags ClassList
// @Summary 分页获取ClassList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ClassList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classlist/getClassListList [get]
export const getClassListList = (params) => {
    return service({
        url: "/classlist/getClassListList",
        method: 'get',
        params
    })
}

export const delAll = (data) => {
    return service({
        url: "/board/delAll",
        method: 'delete',
        data
    })
}