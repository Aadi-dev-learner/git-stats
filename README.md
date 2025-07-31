# git-stats 🧮📊

`git-stats` is a lightweight CLI tool written in **Go** that fetches and displays GitHub repository contribution statistics *without cloning the repo*. It queries the GitHub REST API, aggregates contributor data, and renders it in colorful, easy‑to‑read terminal output.

---

## ✨ Features

- **Zero‑Clone Analysis** — stats pulled straight from the GitHub API.
- **Cross‑Platform** — pre‑built binaries for Linux, macOS (Intel & Apple Silicon), and Windows.
- **One‑Line Install** — curl | bash (Unix) or iwr | iex (Windows).
- **Readable Output** — color‑coded tables showing top contributors, stars, forks, language, and more.

---

## 🚀 Quick Install


### Linux / macOS

```bash
curl -sSf https://raw.githubusercontent.com/Aadi-dev-learner/git-stats/main/scripts/install.sh | sudo bash
```

### Windows (PowerShell)
Make sure to open powershell as adminstrator.
```powershell
iwr -useb https://raw.githubusercontent.com/Aadi-dev-learner/git-stats/main/scripts/install.ps1 | iex
```

*The installer detects your OS/CPU, downloads the latest binary from the ****GitHub Releases**** page, and places it in a directory that’s already on your **`$PATH`** (**`/usr/local/bin`** on Unix or your user folder on Windows).*

---

## ⚡️ Usage

```bash
# Basic usage — show top contributors for a public repo
$ git-stats -r https://github.com/torvalds/linux

Example output:

```
📘 Repo: torvalds/linux
   ⭐ Stars  : 165 k    🍴 Forks  : 55 k    💻 Language : C

👥 Top Contributors (top 10)
 1. torvalds            – 10 231 commits
 2. gregkh              –  8 412 commits
 3. davem               –  6 521 commits
  …
```

---

## 🛠 Building from Source

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/git-stats.git
cd git-stats
go build -o git-stats
```

Cross‑compile binaries (examples):

```bash
GOOS=linux   GOARCH=amd64 go build -o bin/git-stats-linux-amd64
GOOS=darwin  GOARCH=arm64 go build -o bin/git-stats-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o bin/git-stats-windows-amd64.exe
```

---

## 🧑‍💻 Contributing

1. Fork the repository → create your branch → commit changes → open a PR.
2. Run `go test ./...` before submitting (tests coming soon).

---

## 📜 License

`git-stats` is distributed under the **MIT License**. See `LICENSE` for details.

