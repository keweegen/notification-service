package http

import (
    "github.com/gofiber/fiber/v2"
    "github.com/keweegen/notification/internal/entity"
    "github.com/keweegen/notification/internal/service"
)

type userHandler struct {
    services *service.Store
}

func (h *userHandler) init(services *service.Store) *userHandler {
    h.services = services
    return h
}

func (h *userHandler) UpdateChannels(c *fiber.Ctx) error {
    requestData := new(entity.UserChannels)

    if err := c.BodyParser(&requestData); err != nil {
        return sendBadRequest(c, err)
    }
    if err := h.services.User.UpdateNotificationChannels(c.Context(), *requestData); err != nil {
        return sendError(c, err)
    }

    return sendSuccess(c, operationStatus{
        Status:            true,
        StatusDescription: "User notification channels updated",
    })
}
