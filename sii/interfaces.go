package sii

import "github.com/gofiber/fiber/v2"

type Controller interface {
	GetContributorInfo(ctx *fiber.Ctx) error
}

type Cases interface {
	GetContributorInfo(user *User) (ContributorInfo, error)
}

type Repository interface {
	GetContributorInfo(user *User) (ContributorInfo, error)
}
