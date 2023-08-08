package sii

type ContributorInfo struct {
	IdentifierType                       string               `json:"identifier_type"`
	IdentifierNumber                     string               `json:"identifier_number"`
	VerificationCode                     string               `json:"verification_code"`
	CommerceName                         string               `json:"commerce_name"`
	IsInitiatedActivities                bool                 `json:"is_initiated_activities"`
	IsAvailableToPayTaxInForeignCurrency bool                 `json:"is_available_to_pay_tax_in_foreign_currency"`
	IsSmallerCompany                     bool                 `json:"is_smaller_company"`
	CommercialActivities                 []CommercialActivity `json:"commercial_activities"`
}

type CommercialActivity struct {
	Name          string `json:"name"`
	Code          int    `json:"code"`
	Category      int    `json:"category"`
	IsVATAffected bool   `json:"is_vat_affected"`
}

type User struct {
	IdentifierType   string `json:"identifier_type"`
	IdentifierNumber string `json:"identifier_number"`
	VerificationCode string `json:"verification_code"`
}
