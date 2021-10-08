package models

import (
	"github.com/google/uuid"
)

type UserStore interface {
	GetById(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	DeleteById(id uuid.UUID) error
	DeleteByEmail(email string) error

	GetEmailCode(user *User) (*EmailCode, error)
	DeleteEmailCode(code *EmailCode) error

	GetRefreshTokenByCode(user *User, code string) (*RefreshToken, error)
	GetRefreshTokens(user *User) ([]RefreshToken, error)
	AddRefreshToken(user *User, refreshToken *RefreshToken) error
	RotateRefreshToken(user *User, oldRefreshToken *RefreshToken) (*RefreshToken, error)
	DeleteRefreshToken(refreshToken *RefreshToken) error

	GetPasswordTokenByCode(user *User, code string) (*PasswordToken, error)
	GetPasswordTokens(user *User) ([]PasswordToken, error)
	DeletePasswordToken(token *PasswordToken) error

	GetTwoFATokenByCode(user *User, code string) (*TwoFAToken, error)
	GetTwoFATokens(user *User) ([]TwoFAToken, error)
	DeleteTwoFAToken(token *TwoFAToken) error

	GetRecoveryCodeByCode(user *User, code string) (*RecoveryCode, error)
	GetRecoveryCodes(user *User) ([]RecoveryCode, error)
	NewRecoveryCodes(user *User) ([]RecoveryCode, error)
	DeleteRecoveryCode(code *RecoveryCode) error

	GetConfirmEmailLastSent(email string) (int64, error)
	SetConfirmEmailLastSent(email string, time int64) error
}

type User struct {
	Base
	Name             string
	Email            string `gorm:"unique"`
	PasswordHash     []byte
	ProfilePicture   []byte
	ProfilePictureId uuid.UUID `gorm:"type:uuid"`
	EmailConfirmed   bool
	TwoFaOTPEnabled  bool
	OtpSecret        string
	OtpQrCode        []byte
	EmailCode        EmailCode       `gorm:"constraint:OnDelete:CASCADE"`
	RefreshTokens    []RefreshToken  `gorm:"constraint:OnDelete:CASCADE"`
	PasswordTokens   []PasswordToken `gorm:"constraint:OnDelete:CASCADE"`
	TwoFATokens      []TwoFAToken    `gorm:"constraint:OnDelete:CASCADE"`
	RecoveryCodes    []RecoveryCode  `gorm:"constraint:OnDelete:CASCADE"`
}

type ConfirmEmailLastSent struct {
	Base
	Email    string `gorm:"unique"`
	LastSent int64
}

type EmailCode struct {
	Base
	Code           string
	ExpirationTime int64
	UserId         uuid.UUID `gorm:"type:uuid"`
}

type RefreshToken struct {
	Base
	Code           string
	ExpirationTime int64
	Used           bool
	UserId         uuid.UUID `gorm:"type:uuid"`
}

type PasswordToken struct {
	Base
	Code           string
	ExpirationTime int64
	UserId         uuid.UUID `gorm:"type:uuid"`
}

type TwoFAToken struct {
	Base
	Code           string
	ExpirationTime int64
	UserId         uuid.UUID `gorm:"type:uuid"`
}

type RecoveryCode struct {
	Base
	Code   string
	UserId uuid.UUID `gorm:"type:uuid"`
}