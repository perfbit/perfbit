package services

// Repository represents a GitHub repository
type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

// Branch represents a branch in a GitHub repository
type Branch struct {
	Name string `json:"name"`
}

// Commit represents a commit in a GitHub repository
type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
}
