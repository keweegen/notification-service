package http

import "github.com/gofiber/fiber/v2"

func sendSuccess(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func sendBadRequest(c *fiber.Ctx, err error) error {
	return sendError(c, err, fiber.StatusBadRequest)
}

func sendError(c *fiber.Ctx, err error, statusCode ...int) error {
	httpStatusCode := fiber.StatusInternalServerError

	if len(statusCode) > 0 {
		httpStatusCode = statusCode[0]
	}

	return c.Status(httpStatusCode).JSON(operationStatus{
		Status:            false,
		StatusDescription: err.Error(),
	})
}
