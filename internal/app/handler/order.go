package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nick1729/resp-api-tmpl/internal/app/service/order"
)

type orderResource struct {
	service orderService
}

func (or *orderResource) handleGet(ctx *fiber.Ctx) error {
	var req order.GetReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendString("decoding request")
	}

	resp, respErr := or.service.Get(ctx.Context(), &req)
	if respErr != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = ctx.Status(fiber.StatusOK).JSON(resp)
	if err != nil {
		return ctx.SendString("encoding response")
	}

	return nil
}

func (or *orderResource) handleList(ctx *fiber.Ctx) error {
	return nil
}

func (or *orderResource) handleUpdate(ctx *fiber.Ctx) error {
	return nil
}
