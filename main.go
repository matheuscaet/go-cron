package main

import (
	"fmt"
	"time"

	"github.com/matheuscaet/go-cron/external"
	"github.com/robfig/cron"
)

var ENVIRONMENT string = GetEnvironment()

func main() {
	c := cron.New()

	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		posts := external.GetPosts()
		if len(posts) > 0 {
			for _, post := range posts {
				fmt.Printf("[%s] Hello world! The post id is: %d\n", ENVIRONMENT, post.Id)
			}
			fmt.Printf("[%s] Posts size is: %d\n", ENVIRONMENT, len(posts))
		}
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}
