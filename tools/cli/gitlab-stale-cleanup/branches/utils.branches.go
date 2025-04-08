package branches

import (
	"log"
	"sync"
	"time"
)

// isBranchStale checks if the branch is stale based on the commit date
func isBranchStale(commitDate string, staleDaysThreshold int) bool {
	commitTime, err := time.Parse(time.RFC3339, commitDate)
	if err != nil {
		log.Printf("Invalid commit date format for branch %s: %v", commitDate, err)
		return false
	}

	duration := time.Since(commitTime)
	daysOld := int(duration.Hours() / 24)

	return daysOld >= staleDaysThreshold
}

// ProcessBranches processes each branch and deletes it if necessary
func ProcessBranches(branches []Branch, projectID, privateToken string, staleDaysThreshold int, protectedBranches []string, isDryRun bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, branch := range branches {
		// Skip protected branches
		if contains(protectedBranches, branch.Name) {
			log.Printf("Branch '%s' is protected, skipping deletion.", branch.Name)
			continue
		}

		// Check if the branch is stale
		if isBranchStale(branch.Commit.Date, staleDaysThreshold) {
			log.Printf("Branch '%s' is stale, preparing to delete.", branch.Name)

			// Delete the branch if it's stale
			err := DeleteBranch(projectID, privateToken, branch.Name, isDryRun)
			if err != nil {
				log.Printf("Error deleting branch '%s': %v", branch.Name, err)
			}
		}
	}
}

// contains checks if a branch is in the protected branches list
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
