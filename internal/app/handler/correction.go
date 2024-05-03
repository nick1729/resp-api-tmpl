package handler

import (
	"os"

	"github.com/davecgh/go-spew/spew"
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

	b := ctx.Body()

	err := req.UnmarshalJSON(b)

	spew.Dump(req)
	os.Exit(1)

	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendString("decoding request")
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
