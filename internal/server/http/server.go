package http

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/keweegen/notification/internal/service"
)

func NewServer(services *service.Store, addr string) error {
    app := fiber.New(fiber.Config{
        AppName:           "Notification Service",
        EnablePrintRoutes: true,
    })

    app.Use(logger.New())

    setRoutes(services, app)

    return app.Listen(addr)
}

func setRoutes(services *service.Store, app *fiber.App) {
    messageGroup := app.Group("message")
    messageHandlers := new(messageHandler).init(services)
    messageGroup.Post("generate-id", messageHandlers.GenerateID).Name("Generate message id")
    messageGroup.Post(":messageId/send", messageHandlers.Send).Name("Send message by generated id")
    messageGroup.Get(":messageId/status", messageHandlers.GetStatus).Name("Get message status by generated id")
    // -
    userGroup := app.Group("user")
    userHandlers := new(userHandler).init(services)
    userGroup.Patch(":userId/channels", userHandlers.UpdateChannels).Name("CreateOrUpdateChannels user notification channels")
}
