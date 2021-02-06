package request

import "gin-vue-admin/model"

type ClassSearch struct{
    model.Class
    PageInfo
}