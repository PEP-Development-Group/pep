package request

import uuid "github.com/satori/go.uuid"

// Teacher register structure
type RegisterTeacher struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	RealName string `json:"name"`
	// HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`

	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// Student register structure
type RegisterStudent struct {
	Username string `json:"username"`
	Password string `json:"password"`

	College  string `json:"college"`
	Major    string `json:"major"`
	PID      string `json:"pid"`

	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
