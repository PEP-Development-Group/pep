package request

import "pep/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}