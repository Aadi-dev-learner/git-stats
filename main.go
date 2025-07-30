package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Aadi-dev-learner/git-stats/github"
	"github.com/Aadi-dev-learner/git-stats/utils"
)

func main() {
	url := flag.String("r", "", "Github repository link : git-stats -r X (X === repo link)")
	flag.Parse()

	if *url == "" {
		fmt.Println("Please provide a repository url to continue.")
		os.Exit(1)
	}

	repoName := make([]byte, 0)

	for i := 19; i < len(*url); i++ {
		repoName = append(repoName, (*url)[i])
	}

	repo := string(repoName)

	utils.PrintTitle("Fetching data for " + repo)

	contributors, repoInfo, err := github.FetchRepoStats(repo)
	if err != nil {
		utils.PrintError("Failed to fetch stats:" + err.Error())
		os.Exit(1)
	}

	utils.PrintRepoInfo(repoInfo)
	utils.PrintContributors(contributors)
}
