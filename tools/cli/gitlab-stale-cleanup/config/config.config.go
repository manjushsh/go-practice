package config

import (
	"github.com/urfave/cli/v2"
)

// GetFlags returns a slice of flags for the CLI app.
func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "project-id",
			Usage:    "GitLab project ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "private-token",
			Usage:    "Private token for authentication",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "stale-days-threshold",
			Usage:    "Days threshold to consider a branch stale",
			Value:    30,
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "dry-run",
			Usage:    "Dry run mode to preview deletions",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "protected-branches",
			Usage:    "Comma or space-separated list of protected branches",
			Required: false,
		},
	}
}
