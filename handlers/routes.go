package handlers

import (
	"github.com/Bananenpro/hbank2-api/router/middlewares"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterV1(v1 *echo.Group) {
	v1.GET("/status", h.Status)

	auth := v1.Group("/auth")
	auth.POST("/register", h.Register)
	auth.GET("/confirmEmail/:email", h.SendConfirmEmail)
	auth.POST("/confirmEmail", h.VerifyConfirmEmailCode)
	auth.POST("/passwordAuth", h.PasswordAuth)
	auth.POST("/login", h.Login)

	twoFactor := auth.Group("/twoFactor")
	twoFactor.POST("/otp/activate", h.Activate2FAOTP)
	twoFactor.POST("/otp/verify", h.VerifyOTPCode)
	twoFactor.POST("/recovery/get", h.GetRecoveryCodes, middlewares.JWT)
	twoFactor.POST("/recovery/verify", h.VerifyRecoveryCode)
	twoFactor.POST("/recovery/new", h.NewRecoveryCodes, middlewares.JWT)

	user := v1.Group("/user", middlewares.JWT)
	user.GET("/:id", h.GetUser)
}