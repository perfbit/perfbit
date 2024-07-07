package services

import (
	"context"
	"log"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func GetRepositories(token string) ([]*github.Repository, error) {
	log.Println("Fetching repositories")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Printf("Error fetching repositories: %v", err)
		return nil, err
	}

	log.Println("Repositories fetched successfully")
	return repos, nil
}

func GetCommits(token, owner, repo, branch string) ([]*github.RepositoryCommit, error) {
	log.Printf("Fetching commits for owner: %s, repo: %s, branch: %s", owner, repo, branch)
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
		log.Printf("Error fetching commits: %v", err)
		return nil, err
	}

	log.Println("Commits fetched successfully")
	return commits, nil
}

func GetBranches(token, owner, repo string) ([]*github.Branch, error) {
	log.Printf("Fetching branches for owner: %s, repo: %s", owner, repo)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		log.Printf("Error fetching branches: %v", err)
		return nil, err
	}

	log.Println("Branches fetched successfully")
	return branches, nil
}
