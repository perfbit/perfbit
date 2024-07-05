// internal/services/github.go
package services

import (
	"context"
	"log"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func GetRepositories(token string) ([]*github.Repository, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatalf("Error fetching repositories: %v", err)
	}

	return repos, err
}

func GetCommits(token, owner, repo, branch string) ([]*github.Commit, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{
		SHA: branch,
	})
	if err != nil {
		log.Fatalf("Error fetching commits: %v", err)
	}

	return commits, err
}