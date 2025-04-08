package branches

import (
	"fmt"
	"log"
	"net/http"
)

// DeleteBranch deletes a branch in the GitLab project
func DeleteBranch(projectID, privateToken, branchName string, isDryRun bool) error {
	if isDryRun {
		log.Printf("[DRY RUN] Branch '%s' would be deleted.", branchName)
		return nil
	}

	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/repository/branches/%s", projectID, branchName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("PRIVATE-TOKEN", privateToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("Branch '%s' deleted successfully.", branchName)
	} else {
		log.Printf("Failed to delete branch '%s'. Status: %s", branchName, resp.Status)
	}

	return nil
}
