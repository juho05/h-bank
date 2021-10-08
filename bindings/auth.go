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

type Activate2FAOTP struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type VerifyOTPCode struct {
	Email   string `json:"email" form:"email"`
	OTPCode string `json:"otp_code" form:"otp_code"`
}

type PasswordAuth struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Login struct {
	Email         string `json:"email" form:"email"`
	PasswordToken string `json:"password_token" form:"password_token"`
	TwoFAToken    string `json:"two_fa_token" form:"two_fa_token"`
}
