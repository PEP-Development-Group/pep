package request

// User register structure
type Register struct {
	Username int64  `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name" gorm:"default:'系统用户'"`

	AuthorityId string `json:"authorityId"`

	Class string `json:"class"`

	CancelNums   int `json:"cancel_nums"`
	TotalCredits int `json:"total_credits"`
}

// User login structure
type Login struct {
	Username  int64  `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    int64  `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UsernameRequest
	AuthorityId string `json:"authorityId"`
}

type AddCancelNums struct {
	Username int64 `json:"username"`
	Cnt      int   `json:"cnt"`
}
