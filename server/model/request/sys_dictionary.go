package request

import "pep/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
