package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
