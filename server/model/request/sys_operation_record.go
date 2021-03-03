package request

import "pep/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
