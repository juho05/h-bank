package bindings

type Register struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	CaptchaToken string `json:"h-captcha-response" form:"h-captcha-response"`
}

type ConfirmEmail struct {
	Email string `json:"email" form:"email"`
	Code  string `json:"code" form:"code"`
}
