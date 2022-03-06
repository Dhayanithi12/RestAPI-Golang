package main

import (
	"fmt"
)
c := cron.New()
c.AddFunc("@every 1min", SendMessage)
c.Start()
func SendMessage() {
fmt.Println("Successfully! Mail has been sent!.")
}
