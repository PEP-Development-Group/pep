package utils

import (
	"github.com/tealeg/xlsx/v3"
)

//@author: [sh1luo](https://github.com/sh1luo)
//@function: ParseFile
//@description: 解析xlsx文件获取记录数组
//@param: bs []byte
//@return: error

func ParseFile(bs []byte) error {
	wb, err := xlsx.OpenBinary(bs)
	if err != nil {
		return err
	}
	sh, exist := wb.Sheet["student"]
	if !exist {
		return err
	}
	return sh.ForEachRow(func(r *xlsx.Row) error {
		return r.ForEachCell(func(c *xlsx.Cell) error {

		})
	})
}