package http

import (
    "github.com/gofiber/fiber/v2"
    "github.com/keweegen/notification/internal/channel"
    "github.com/keweegen/notification/internal/messagetemplate"
    "github.com/keweegen/notification/internal/service"
)

type messageHandler struct {
    services *service.Store
}

func (h *messageHandler) init(services *service.Store) *messageHandler {
    h.services = services
    return h
}

func (h *messageHandler) GenerateID(c *fiber.Ctx) error {
    requestData := new(generateMessageIDRequest)

    if err := c.BodyParser(&requestData); err != nil {
        return sendBadRequest(c, err)
    }

    ch, ok := channel.GetChannelTypeFromString(requestData.Channel)
    if !ok {
        return sendError(c, service.InvalidChannelErr)
    }

    mt, ok := messagetemplate.GetMessageTemplateTypeFromString(requestData.MessageTemplate)
    if !ok {
        return sendError(c, service.InvalidMessageTemplateErr)
    }

    id := h.services.Message.GenerateID(
        ch,
        mt,
        requestData.UserID,
        requestData.Timestamp,
        requestData.ExternalID)

    return sendSuccess(c, generateMessageIDResponse{ID: id})
}

func (h *messageHandler) Send(c *fiber.Ctx) error {
    messageID := c.Params("messageId")
    requestData := new(sendMessageRequest)

    if err := c.BodyParser(&requestData); err != nil {
        return sendBadRequest(c, err)
    }

    messageID, err := h.services.Message.Send(c.Context(), messageID, requestData.Params)
    if err != nil {
        return sendError(c, err)
    }

    status, err := h.services.Message.GetStatus(c.Context(), messageID)
    if err != nil {
        return sendError(c, err)
    }

    return sendSuccess(c, messageResponse{
        ID:                status.MessageID,
        Status:            status.Status,
        StatusDescription: status.Description,
        StatusTime:        status.CreatedAt,
    })
}

func (h *messageHandler) GetStatus(c *fiber.Ctx) error {
    status, err := h.services.Message.GetStatus(c.Context(), c.Params("messageId"))
    if err != nil {
        return sendError(c, err)
    }

    return sendSuccess(c, messageResponse{
        ID:                status.MessageID,
        Status:            status.Status,
        StatusDescription: status.Description,
        StatusTime:        status.CreatedAt,
    })
}
