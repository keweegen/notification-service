package cmd

import (
    "github.com/keweegen/notification/config"
    "github.com/keweegen/notification/logger"
    "github.com/spf13/cobra"
    "log"
)

var (
    l   logger.Logger
    cfg *config.Config
)

var rootCmd = &cobra.Command{
    Use: "notification-service",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    cobra.OnInitialize(initLogger, initConfig)
}

func initLogger() {
    var err error
    l, err = logger.NewLogger()
    if err != nil {
        log.Fatalln("init logger", err)
    }
}

func initConfig() {
    var err error
    cfg, err = config.Read()
    if err != nil {
        log.Fatalln("init config", err)
    }
}
