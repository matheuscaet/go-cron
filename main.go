package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		posts := [1]int{1024}
		if len(posts) > 0 {
			fmt.Println("Hello world! The first post is:", posts[0])
		}
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}
