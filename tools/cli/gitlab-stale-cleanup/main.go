package main

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/manjushsh/go-practice/tools/gitlab-stale-cleanup/branches"
	"github.com/manjushsh/go-practice/tools/gitlab-stale-cleanup/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gitlab-stale-branch-cleaner",
		Usage: "A tool to delete stale branches in a GitLab repository",
		Flags: config.GetFlags(), // Load flags from config package
		Action: func(c *cli.Context) error {
			// Read flags from CLI input
			projectID := c.String("project-id")
			privateToken := c.String("private-token")
			staleDaysThreshold := c.Int("stale-days-threshold")
			isDryRun := c.Bool("dry-run")
			protectedBranches := strings.Fields(c.String("protected-branches"))

			// Fetch branches
			allBranches, err := branches.FetchBranches(projectID, privateToken)
			if err != nil {
				return err
			}

			// Use WaitGroup for concurrency
			var wg sync.WaitGroup
			wg.Add(1)

			// Process branches concurrently
			go branches.ProcessBranches(allBranches, projectID, privateToken, staleDaysThreshold, protectedBranches, isDryRun, &wg)

			// Wait for all goroutines to finish
			wg.Wait()

			return nil
		},
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
