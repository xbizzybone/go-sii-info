package sii

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/go-sii-info/utils"
	"go.uber.org/zap"
)

var repo Repository
var ctrl Controller
var cs Cases

func bootstrap(logger *zap.Logger) {
	repo = NewRepository(logger)
	cs = NewCases(repo)
	ctrl = NewController(cs)
}

func ApplyRoutes(app *fiber.App, logger *zap.Logger) {
	bootstrap(logger)

	group := app.Group("/sii", utils.GetNextMiddleWare)

	group.Get("/contributor-info/:identifier_type/:identifier_number", ctrl.GetContributorInfo)
}
