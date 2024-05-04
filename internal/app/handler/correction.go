package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nick1729/resp-api-tmpl/internal/app/service/correction"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/http/protocol"
)

type correctionResource struct {
	service correctionService
}

func (cr *correctionResource) handleGet(ctx *fiber.Ctx) error {
	var (
		req  protocol.Request
		data correction.GetReq
	)

	err := req.UnmarshalJSON(ctx.Body())
	if err != nil {
		return ctx.SendString("decoding request")
	}

	err = data.UnmarshalJSON(req.Data)
	if err != nil {
		return ctx.SendString("decoding data")
	}

	resp, err := cr.service.Get(ctx.Context(), &data)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = ctx.Status(fiber.StatusOK).JSON(resp)
	if err != nil {
		return ctx.SendString("encoding response")
	}

	return nil
}

func (cr *correctionResource) handleList(ctx *fiber.Ctx) error {
	return nil
}

func (cr *correctionResource) handleUpdate(ctx *fiber.Ctx) error {
	return nil
}
