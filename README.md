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


### Linux

```bash
curl -LO https://github.com/jain-aadi/git-stats/releases/download/v1.0.0/git-stats-linux-amd64
chmod +x git-stats-linux-amd64
sudo mv git-stats-linux-amd64 /usr/local/bin/git-stats

```
---
### MacOS (M1/M2...)

```bash
curl -LO https://github.com/jain-aadi/git-stats/releases/download/v1.0.0/git-stats-darwin-arm64
chmod +x git-stats-darwin-arm64
sudo mv git-stats-darwin-arm64 /usr/local/bin/git-stats

```
### MacOS (Intel)

```bash
curl -LO https://github.com/jain-aadi/git-stats/releases/download/v1.0.0/git-stats-darwin-amd64
chmod +x git-stats-darwin-amd64
sudo mv git-stats-darwin-amd64 /usr/local/bin/git-stats

```
### Windows (powershell : run as adminstrator)

```powershell 
iwr -useb https://github.com/jain-aadi/git-stats/releases/download/v1.0.0/git-stats-windows-amd64.exe -OutFile $env:USERPROFILE\git-stats.exe
$env:Path += ";$env:USERPROFILE"

```


-- Use git-stats.exe for windows .
## ⚡️ Usage 

```bash
# Basic usage — show top contributors for a public repo
$ git-stats -r https://github.com/torvalds/linux

Example output:

📘 Repo: torvalds/linux
   ⭐ Stars  : 165 k    🍴 Forks  : 55 k    💻 Language : C

👥 Top Contributors (top 10)
 1. torvalds            – 10 231 commits
 2. gregkh              –  8 412 commits
 3. davem               –  6 521 commits
  …
```

---

## 🧑‍💻 Contributing

1. Fork the repository → create your branch → commit changes → open a PR.
2. Run `go test ./...` before submitting (tests coming soon).

---

## 📜 License

`git-stats` is distributed under the **MIT License**. See `LICENSE` for details.

