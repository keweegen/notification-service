package main

import (
    "github.com/keweegen/notification/cmd"
    "log"
)

func main() {
    if err := cmd.Execute(); err != nil {
        log.Fatalln("failed to cmd execute", err)
    }
}
