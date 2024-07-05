// internal/services/metrics.go
package services

func GeneratePerformanceMetrics(commits []*github.Commit) map[string]map[string]int {
	metrics := make(map[string]map[string]int)
	for _, commit := range commits {
		author := *commit.Author.Login
		if metrics[author] == nil {
			metrics[author] = make(map[string]int)
		}
		metrics[author]["commits"]++
		// Add other metrics as needed
	}
	return metrics
}
