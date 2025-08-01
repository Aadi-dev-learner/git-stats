package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jain-aadi/git-stats/github"
)

func PrintTitle(title string) {
	color.New(color.FgCyan, color.Bold).Println("\n==> " + title)
}

func PrintError(err string) {
	color.New(color.FgRed, color.Bold).Println("Error: " + err)
}

func PrintRepoInfo(info github.RepoInfo) {
	color.Green("\n📘 Repo: %s\n", info.FullName)
	fmt.Printf("   📄 %s\n", info.Description)
	fmt.Printf("   ⭐ Stars: %d\n", info.StargazersCount)
	fmt.Printf("   🍴 Forks: %d\n", info.ForksCount)
	fmt.Printf("   💻 Language: %s\n", info.Language)
}

func PrintContributors(contributors []github.Contributor) {
	color.Blue("\n👥 Top Contributors:\n")
	for i, c := range contributors {
		fmt.Printf("   %2d. %-20s 📝 %d commits\n", i+1, c.Login, c.Contributions)
		if i == 9 {
			break
		}
	}
}
