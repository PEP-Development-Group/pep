package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
