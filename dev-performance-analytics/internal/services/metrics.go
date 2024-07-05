// internal/services/metrics.go
package services

import "github.com/google/go-github/v39/github"

func GeneratePerformanceMetrics(commits []*github.RepositoryCommit) map[string]map[string]int {
	metrics := make(map[string]map[string]int)
	for _, commit := range commits {
		author := *commit.Commit.Author.Name
		if metrics[author] == nil {
			metrics[author] = make(map[string]int)
		}
		metrics[author]["commits"]++
		// Add other metrics as needed
	}
	return metrics
}
