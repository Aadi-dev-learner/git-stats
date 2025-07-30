package github

type Contributor struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

type RepoInfo struct {
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	Language        string `json:"language"`
}
