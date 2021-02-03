package utils

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
)

// 姓名，学号，学院，专业，身份证号
type Student struct {
	Name      string
	StuNumber string
	College   string
	Major     string
	PID       string
}

func ParseExcelFile(bs string) (*[]Student, error) {
	wb, err := xlsx.OpenFile(bs)
	if err != nil {
		return nil, err
	}
	sh, exist := wb.Sheet["Sheet1"]
	if !exist {
		return nil, errors.New("sheet not exist")
	}
	var st []Student
	err = sh.ForEachRow(func(r *xlsx.Row) error {
		var s Student
		s.Name = r.GetCell(0).String()
		s.StuNumber = r.GetCell(1).String()
		s.College = r.GetCell(2).String()
		s.Major = r.GetCell(3).String()
		s.PID = r.GetCell(4).String()
		st = append(st, s)
		return nil
	})
	return &st, err
}
