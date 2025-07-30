package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.github.com/repos/"

func FetchRepoStats(repo string) ([]Contributor, RepoInfo, error) {
	contriburl := fmt.Sprintf("%s%s/contributors", baseURL, repo)
	repourl := fmt.Sprintf("%s%s", baseURL, repo)

	contributors := []Contributor{}
	repoinfo := RepoInfo{}

	if err := getJSON(contriburl, &contributors); err != nil {
		return nil, repoinfo, err
	}

	if err := getJSON(repourl, &repoinfo); err != nil {
		return nil, repoinfo, err
	}

	return contributors, repoinfo, nil
}

func getJSON(url string, target interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "git-stats-cli")

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("Github API error: %s", response.Status)
	}

	return json.NewDecoder(response.Body).Decode(target)
}
