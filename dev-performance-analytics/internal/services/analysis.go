// internal/services/analysis.go
package services

import "github.com/google/go-github/v39/github"

func AnalyzeCommits(commits []*github.Commit) map[string]int {
	contributions := make(map[string]int)
	for _, commit := range commits {
		author := *commit.Author.Login
		contributions[author]++
	}
	return contributions
}
