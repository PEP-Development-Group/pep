package request

import "gin-vue-admin/model"

type ClassListSearch struct{
    model.ClassList
    PageInfo
}