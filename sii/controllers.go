package sii

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/go-sii-info/utils"
)

type controller struct {
	cases Cases
}

func NewController(cases Cases) Controller {
	return &controller{
		cases: cases,
	}
}

func (c *controller) GetContributorInfo(ctx *fiber.Ctx) error {
	userRequest := UserRequest{
		IdentifierType:   ctx.Params("identifier_type"),
		IdentifierNumber: ctx.Params("identifier_number"),
	}

	if !utils.IsValidRut(userRequest.IdentifierNumber) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Invalid identifier number: %s", userRequest.IdentifierNumber),
		})
	}

	user := User{}

	switch userRequest.IdentifierType {
	case "RUN", "RUT":
		formatRut := utils.Format(userRequest.IdentifierNumber)

		user = User{
			IdentifierType:   userRequest.IdentifierType,
			IdentifierNumber: utils.GetRutWithoutVerificationCode(formatRut),
			VerificationCode: utils.GetVerificationCode(formatRut),
		}
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Invalid identifier type: %s", userRequest.IdentifierType),
		})
	}

	result, err := c.cases.GetContributorInfo(&user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error fetching contributor info: %s", err.Error()),
		})
	}

	contributorInfoResponse := ContributorInfoResponse{
		IdentifierType:                       result.IdentifierType,
		IdentifierNumber:                     result.IdentifierNumber,
		VerificationCode:                     result.VerificationCode,
		CommerceName:                         result.CommerceName,
		IsInitiatedActivities:                result.IsInitiatedActivities,
		IsAvailableToPayTaxInForeignCurrency: result.IsAvailableToPayTaxInForeignCurrency,
		IsSmallerCompany:                     result.IsSmallerCompany,
		CommercialActivities:                 result.CommercialActivities,
		StampedDocuments:                     result.StampedDocuments,
	}

	return ctx.JSON(contributorInfoResponse)
}
