package constant

import "errors"

var (
	ErrClassNotExist = errors.New("课程不存在")
	ErrClassHasFull  = errors.New("选课人数已满")
)
