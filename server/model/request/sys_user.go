package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name" gorm:"default:'用户'"`

	// HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`

	AuthorityId string `json:"authorityId"`

	College string `json:"college"`
	Major   string `json:"major"`
	PID     string `json:"pid"`
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
