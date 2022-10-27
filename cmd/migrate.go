package cmd

import (
    "database/sql"
    "fmt"
    "github.com/keweegen/notification/db"
    "github.com/keweegen/notification/internal/app"
    "github.com/pressly/goose/v3"
    "github.com/spf13/cobra"
)

var migrateDir = "migrations"

func init() {
    rootCmd.AddCommand(migrateCommand)
}

var migrateCommand = &cobra.Command{
    Use:     "migrate",
    Args:    cobra.MinimumNArgs(1),
    Example: "migrate up",
    RunE: func(_ *cobra.Command, args []string) error {
        app := app.New(cfg, l)
        if err := app.OpenConnectDatabase(); err != nil {
            return err
        }
        defer func() {
            if err := app.CloseConnectDatabase(); err != nil {
                app.Logger.Error("failed to close connection database", "error", err)
            }
        }()

        return runMigrate(app.CurrentDatabase(), migrateDir, args[0])
    },
}

func runMigrate(database *sql.DB, dir string, action string) error {
    goose.SetBaseFS(db.EmbedMigrations)
    goose.SetTableName("migrations")
    if err := goose.SetDialect("postgres"); err != nil {
        return err
    }

    switch action {
    case "up":
        return goose.Up(database, dir)
    case "down":
        return goose.Down(database, dir)
    case "status":
        return goose.Status(database, dir)
    case "redo":
        return goose.Redo(database, dir)
    case "reset":
        return goose.Reset(database, dir)
    default:
        return fmt.Errorf("invalid migration action '%s'", action)
    }
}
