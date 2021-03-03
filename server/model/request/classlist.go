package request

import "pep/model"

type ClassListSearch struct {
	model.ClassList
	PageInfo
}
