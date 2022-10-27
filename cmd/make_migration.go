package cmd

import (
    "github.com/pressly/goose/v3"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(makeMigrationCommand)
}

var makeMigrationCommand = &cobra.Command{
    Use:     "make:migration",
    Args:    cobra.MinimumNArgs(1),
    Example: "make:migration create_message_table",
    RunE: func(_ *cobra.Command, args []string) error {
        return makeMigration("db/"+migrateDir, args[0])
    },
}

func makeMigration(dir string, name string) error {
    return goose.Create(nil, dir, name, "sql")
}
