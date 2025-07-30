#!/bin/bash

set -e

ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

# Normalize architecture names
if [[ "$ARCH" == "x86_64" ]]; then ARCH="amd64"; fi
if [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then ARCH="arm64"; fi

BINARY_NAME="git-stats-${OS}-${ARCH}"
URL="https://github.com/Aadi-dev-learner/git-stats/releases/latest/download/${BINARY_NAME}"

echo "⬇️ Downloading git-stats for $OS/$ARCH..."
curl -L "$URL" -o /usr/local/bin/git-stats
chmod +x /usr/local/bin/git-stats

echo "✅ Installed git-stats!"
echo "To run: git-stats -r <repo_link>"
