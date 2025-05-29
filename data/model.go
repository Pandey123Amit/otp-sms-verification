package data

type OTPdata struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type Verifydata struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"phoneNumber,omitempty" validate:"required"`
}
