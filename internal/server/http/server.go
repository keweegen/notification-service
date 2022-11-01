package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/keweegen/notification/internal/service"
)

type Server struct {
	base *fiber.App
}

func NewServer(services *service.Store) *Server {
	server := new(Server)
	server.base = fiber.New(fiber.Config{
		AppName:           "Notification Service",
		EnablePrintRoutes: true,
	})

	server.base.Use(logger.New())
	server.setRoutes(services)

	return server
}

func (s *Server) setRoutes(services *service.Store) {
	messageGroup := s.base.Group("message")
	messageHandlers := new(messageHandler).init(services)
	messageGroup.Post("generate-id", messageHandlers.GenerateID).Name("Generate message id")
	messageGroup.Post(":messageId/send", messageHandlers.Send).Name("Send message by generated id")
	messageGroup.Get(":messageId/status", messageHandlers.GetStatus).Name("Get message status by generated id")

	userGroup := s.base.Group("user")
	userHandlers := new(userHandler).init(services)
	userGroup.Get("channel/:userChannelId", userHandlers.ReadChannel).Name("Get user notification channel")
	userGroup.Get("channel/:userId/all", userHandlers.AllChannelsByUser).Name("Get all notification channels by user id")
	userGroup.Post("channel", userHandlers.CreateChannel).Name("Create user notification channel")
	userGroup.Patch("channel/:userChannelId", userHandlers.UpdateChannel).Name("Update user notification channel")
	userGroup.Delete("channel/:userChannelId", userHandlers.DestroyChannel).Name("Destroy user notification channel")
}

func (s *Server) Listen(addr string) error {
	return s.base.Listen(addr)
}

func (s *Server) Close() error {
	return s.base.Shutdown()
}
