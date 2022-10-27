package cmd

import (
    "context"
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
        s.Listen()

        quit := make(chan bool)
        s.AddHandler("quit handle messages", func() error {
            quit <- true
            return nil
        })

        app := app.New(cfg, l)
        if err := app.Connect(); err != nil {
            return fmt.Errorf("app connect: %w", err)
        }
        defer func() {
            if err := app.Close(); err != nil {
                app.Logger.Error("failed to close app", "error", err)
            }
        }()

        repositoryStore := repository.NewStore(app.CurrentDatabase(), app.CurrentMessageBroker())
        channelStore := channel.NewStore(cfg.NotificationChannels)
        serviceStore := service.NewStore(l, repositoryStore, channelStore)

        go serviceStore.Message.HandleMessages(context.Background(), quit)

        if err := http.NewServer(serviceStore, httpAddr); err != nil {
            return err
        }

        s.ReadCh()

        return nil
    },
}
