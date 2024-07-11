
package main

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"
)
c := cron.New()
c.AddFunc("@every 1m", SendMessage)
c.Start()
func SendMessage() {
fmt.Println("Successfully! Mail has been sent!.")
}
