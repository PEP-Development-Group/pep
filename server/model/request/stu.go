package request

import "gin-vue-admin/model"

type StuSearch struct{
    model.Stu
    PageInfo
}