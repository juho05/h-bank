package bindings

type CreateGroup struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	OnlyAdmin   bool   `json:"onlyAdmin" form:"onlyAdmin"`
}

type UpdateGroup struct {
	Description string `json:"description" from:"description"`
}

type CreateTransaction struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Amount      uint   `json:"amount" form:"amount"`
	ReceiverId  string `json:"receiverId" form:"receiverId"`
	FromBank    bool   `json:"fromBank" form:"fromBank"`
}

type CreatePaymentPlan struct {
	Name         string `json:"name" form:"name"`
	Description  string `json:"description" form:"description"`
	Amount       uint   `json:"amount" form:"amount"`
	ReceiverId   string `json:"receiverId" form:"receiverId"`
	FromBank     bool   `json:"fromBank" form:"fromBank"`
	Schedule     uint   `json:"schedule" form:"schedule"`
	ScheduleUnit string `json:"scheduleUnit" form:"scheduleUnit"`
	// UTC date of first payment with format "YYYY-MM-DD"
	FirstPayment string `json:"firstPayment"`
	// negative payment count for unlimited payments
	PaymentCount int `json:"paymentCount"`
}

type UpdatePaymentPlan struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Amount      uint   `json:"amount" form:"amount"`
	// UTC date of next payment with format "YYYY-MM-DD"
	NextPayment  string `json:"nextPayment"`
	Schedule     uint   `json:"schedule" form:"schedule"`
	ScheduleUnit string `json:"scheduleUnit" form:"scheduleUnit"`
}

type CreateInvitation struct {
	Message string `json:"message" form:"message"`
	UserId  string `json:"userId" form:"userId"`
}
