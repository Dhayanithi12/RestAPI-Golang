package main

import (
	"fmt"
)
c := cron.New()
c.AddFunc("@every 1m", SendMessage)
c.Start()
func SendMessage() {
fmt.Println("Successfully! Mail has been sent!.")
}
