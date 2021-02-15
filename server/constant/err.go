package constant

import "errors"

var (
	ErrClassNotExist = errors.New("课程不存在")

	ErrClassHasFull        = errors.New("选课人数已满")
	ErrClassHasSelected    = errors.New("课程已选")
	ErrClassHasNotSelected = errors.New("未选择该课程")

	ErrDelClassTooMany      = errors.New("退课次数太多")
	ErrDelClassOnDayOfClass = errors.New("上课当天无法退课")
	ErrDelClass             = errors.New("退课失败")
)
