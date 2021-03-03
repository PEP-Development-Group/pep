package request

import "pep/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
