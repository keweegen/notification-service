package cmd

import (
	"fmt"
	"github.com/keweegen/notification/internal/app"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/internal/server/http"
	"github.com/keweegen/notification/internal/service"
	"github.com/keweegen/notification/shutdown"
	"github.com/spf13/cobra"
)

var httpAddr string

func init() {
	httpServerCommand.Flags().StringVarP(&httpAddr, "addr", "a", ":3000", "")
	rootCmd.AddCommand(httpServerCommand)
}

var httpServerCommand = &cobra.Command{
	Use: "http",
	RunE: func(_ *cobra.Command, _ []string) error {
		s := shutdown.New(l)
		defer s.ListenContextDone()

		app := app.New(cfg, l)

		if err := app.Connect(); err != nil {
			return fmt.Errorf("app connect: %w", err)
		}
		s.AddHandler("app", app.Close)

		repositoryStore := repository.NewStore(app.CurrentDatabase(), app.CurrentMessageBroker())
		channelStore := channel.NewStore(cfg.NotificationChannels)
		serviceStore := service.NewStore(l, repositoryStore, channelStore)

		httpServer := http.NewServer(serviceStore)
		s.AddHandler("httpServer", httpServer.Close)

		go serviceStore.Message.HandleMessages(s.Context())
		go serviceStore.MessageChecker.Do(s.Context())
		go func() {
			if err := httpServer.Listen(httpAddr); err != nil {
				l.Error("failed to listen http server", "error", err)
			}
		}()

		return nil
	},
}
