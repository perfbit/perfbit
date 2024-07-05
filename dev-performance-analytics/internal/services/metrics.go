// internal/services/metrics.go
package services

import (
	"strings"

	"github.com/google/go-github/v39/github"
)

func GeneratePerformanceMetrics(commits []*github.RepositoryCommit) map[string]map[string]int {
	metrics := make(map[string]map[string]int)
	for _, commit := range commits {
		author := *commit.Author.Login
		if metrics[author] == nil {
			metrics[author] = make(map[string]int)
		}

		// Initialize metrics
		metrics[author]["commits"]++
		linesAdded, linesDeleted, filesChanged := getCommitStats(commit)
		metrics[author]["lines_added"] += linesAdded
		metrics[author]["lines_deleted"] += linesDeleted
		metrics[author]["files_changed"] += filesChanged

		// Example: Count commits with "fix" in the commit message
		if strings.Contains(strings.ToLower(commit.GetCommit().GetMessage()), "fix") {
			metrics[author]["fix_commits"]++
		}
	}
	return metrics
}

func getCommitStats(commit *github.RepositoryCommit) (int, int, int) {
	stats := commit.GetStats()
	if stats == nil {
		return 0, 0, 0
	}
	return stats.GetAdditions(), stats.GetDeletions(), len(commit.Files)
}
