package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keweegen/notification/internal/channel"
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

func (h *userHandler) CreateChannel(c *fiber.Ctx) error {
	requestData := new(userChannelRequest)
	if err := c.BodyParser(&requestData); err != nil {
		return sendBadRequest(c, err)
	}

	userChannel, err := h.userChannelRequestToEntity(requestData)
	if err != nil {
		return sendError(c, err)
	}
	if err = h.services.User.CreateNotificationChannel(c.Context(), userChannel); err != nil {
		return sendError(c, err)
	}

	return sendSuccess(c, operationStatus{
		Status:            true,
		StatusDescription: "User notification channel created successfully",
	})
}

func (h *userHandler) ReadChannel(c *fiber.Ctx) error {
	channelID, err := c.ParamsInt("userChannelId")
	if err != nil {
		return sendError(c, err)
	}

	ch, err := h.services.User.FindNotificationChannel(c.Context(), int64(channelID))
	if err != nil {
		return sendError(c, err)
	}

	return sendSuccess(c, h.userChannelToResponse(ch))
}

func (h *userHandler) UpdateChannel(c *fiber.Ctx) error {
	channelID, err := c.ParamsInt("userChannelId")
	if err != nil {
		return sendError(c, err)
	}

	requestData := new(userChannelUpdateRequest)
	if err = c.BodyParser(&requestData); err != nil {
		return sendBadRequest(c, err)
	}

	err = h.services.
		User.
		UpdateNotificationChannel(c.Context(), int64(channelID), requestData.Recipient, requestData.CanNotify)
	if err != nil {
		return sendError(c, err)
	}

	return sendSuccess(c, operationStatus{
		Status:            true,
		StatusDescription: "User notification channel updated successfully",
	})
}

func (h *userHandler) DestroyChannel(c *fiber.Ctx) error {
	channelID, err := c.ParamsInt("userChannelId")
	if err != nil {
		return sendError(c, err)
	}
	if err = h.services.User.DestroyNotificationChannel(c.Context(), int64(channelID)); err != nil {
		return sendError(c, err)
	}

	return sendSuccess(c, operationStatus{
		Status:            true,
		StatusDescription: "User notification channel destroyed successfully",
	})
}

func (h *userHandler) AllChannelsByUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userId")
	if err != nil {
		return sendError(c, err)
	}

	channels, err := h.services.User.FindNotificationChannels(c.Context(), int64(userID))
	if err != nil {
		return sendError(c, err)
	}

	return sendSuccess(c, h.userChannelsToResponse(channels))
}

// -- Helpers

func (h *userHandler) userChannelToResponse(channel *entity.UserChannel) *userChannelResponse {
	return &userChannelResponse{
		ID:        channel.ID,
		UserID:    channel.UserID,
		Channel:   channel.Channel.String(),
		Recipient: channel.Recipient,
		CanNotify: channel.CanNotify,
	}
}

func (h *userHandler) userChannelRequestToEntity(channelRequest *userChannelRequest) (*entity.UserChannel, error) {
	ch, ok := channel.GetChannelTypeFromString(channelRequest.Channel)
	if !ok {
		return nil, service.InvalidChannelErr
	}

	return &entity.UserChannel{
		ID:        channelRequest.ID,
		UserID:    channelRequest.UserID,
		Channel:   ch,
		Recipient: channelRequest.Recipient,
		CanNotify: channelRequest.CanNotify,
	}, nil
}

func (h *userHandler) userChannelsToResponse(channels entity.UserChannels) []*userChannelResponse {
	response := make([]*userChannelResponse, 0, len(channels))

	for _, c := range channels {
		response = append(response, h.userChannelToResponse(c))
	}

	return response
}
