package sii

//Request
type UserRequest struct {
	IdentifierType   string `json:"identifier_type" validate:"required,oneof=RUT RUN"`
	IdentifierNumber string `json:"identifier_number" validate:"required"`
}

//Response
type ContributorInfoResponse struct {
	IdentifierType                       string               `json:"identifier_type"`
	IdentifierNumber                     string               `json:"identifier_number"`
	VerificationCode                     string               `json:"verification_code"`
	CommerceName                         string               `json:"commerce_name"`
	IsInitiatedActivities                bool                 `json:"is_initiated_activities"`
	IsAvailableToPayTaxInForeignCurrency bool                 `json:"is_available_to_pay_tax_in_foreign_currency"`
	IsSmallerCompany                     bool                 `json:"is_smaller_company"`
	CommercialActivities                 []CommercialActivity `json:"commercial_activities"`
}

type CaptchaResponse struct {
	TxtCaptcha string `json:"txtCaptcha"`
}
