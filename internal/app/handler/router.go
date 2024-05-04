package handler

import (
	"github.com/gofiber/fiber/v2"

	repo "github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
)

const apiGroupPrefix = "resp-api-tmpl/v1"

func RouteRegister(app *fiber.App, repo *repo.Repository) {
	res := newResources(repo)

	apiGroup := app.Group(apiGroupPrefix)

	apiGroup.Get("/healthcheck", healthcheck)

	correctionGroup := apiGroup.Group("/correction")
	correctionGroup.Post("/get", res.correctionResource.handleGet)
	correctionGroup.Post("/list", res.correctionResource.handleList)
	correctionGroup.Post("/update", res.correctionResource.handleUpdate)

	orderGroup := apiGroup.Group("/order")
	orderGroup.Post("/get", res.orderResource.handleGet)
	orderGroup.Post("/list", res.orderResource.handleList)
	orderGroup.Post("/update", res.orderResource.handleUpdate)
}
