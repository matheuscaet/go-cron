package main

import (
	"fmt"
	"time"

	"github.com/matheuscaet/go-cron/external"
	"github.com/matheuscaet/go-cron/logger"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		posts := external.GetPosts()
		if len(posts) > 0 {
			for _, post := range posts {
				logger.Info("Post ID: " + fmt.Sprint(post.Id))
			}
			logger.Info("Posts size is: " + fmt.Sprint(len(posts)))
		}
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}
