package service

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
	"mime/multipart"
	"os"
	"pep/global"
	"pep/model"
	"pep/model/request"
	"pep/utils"
	"pep/utils/upload"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func Upload(file model.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.ExaFileUploadAndDownload

func FindFile(id uint) (error, model.ExaFileUploadAndDownload) {
	var file model.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func DeleteFile(file model.ExaFileUploadAndDownload) (err error) {
	var fileFromDb model.ExaFileUploadAndDownload
	err, fileFromDb = FindFile(file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []model.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: UploadFile
//@description: 上传xlsx文件，批量导入学生信息
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func UploadFile(header *multipart.FileHeader, mayConflict string) (err error, file model.ExaFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	s := strings.Split(header.Filename, ".")
	f := model.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if err = parseAndCreate(filePath, mayConflict); err != nil {
		return err, f
	}
	return nil, f
}

func parseAndCreate(filename string, mayConflict string) error {
	st, err := ParseExcelFile(filename)
	if err != nil {
		return err
	}
	if mayConflict == "1" {
		for _, s := range *st {
			global.GVA_DB.Select("name").Where(model.SysUser{Username: s.Username}).Assign(model.SysUser{Class: s.Class,
				Name: s.Name,
				TotalCredits: s.TotalCredits,
				Password: s.Password,
				AuthorityId: s.AuthorityId,
				CancelNums: global.GVA_CONFIG.System.CancelNums,
			}).FirstOrCreate(&s)
		}
		return nil
	}
	if err = global.GVA_DB.Create(st).Error; err != nil {
		return err
	}
	if err = os.Remove(filename); err != nil {
		return err
	}
	return nil
}

func ParseExcelFile(bs string) (*[]model.SysUser, error) {
	wb, err := xlsx.OpenFile(bs)
	if err != nil {
		return nil, err
	}
	sh, exist := wb.Sheet["Sheet1"]
	if !exist {
		return nil, errors.New("sheet not exist")
	}
	var st []model.SysUser
	err = sh.ForEachRow(func(r *xlsx.Row) error {
		var s model.SysUser
		s.Class = r.GetCell(0).String()
		s.Username, _ = r.GetCell(1).Int64()
		s.Name = r.GetCell(2).String()
		s.TotalCredits, _ = r.GetCell(4).Int()
		s.Password = utils.MD5V([]byte(utils.MD5V([]byte(r.GetCell(3).String())))) // 密码身份证后8位，两次md5
		s.AuthorityId = "1"

		s.CancelNums = global.GVA_CONFIG.System.CancelNums
		st = append(st, s)
		return nil
	})
	st = st[1:] // 去掉表头
	return &st, err
}
