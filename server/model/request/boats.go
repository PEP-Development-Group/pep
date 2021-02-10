package request

import "gin-vue-admin/model"

type BoatsSearch struct {
	model.Boats
	PageInfo
}
