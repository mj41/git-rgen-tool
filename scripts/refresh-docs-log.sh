#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

set -e

echo "Source repository gl-exA-src latest commit:"
cd ../gl-exA-src
git log --oneline -n1
echo

echo "Repository generation:"
cd ../git-rgen-tool
echo "Removing existing gl-exA directory and regenerating with git-rgen..."
echo "Commands:"
echo "  rm -rf ../gl-exA"
echo "  go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json"
echo
rm -rf ../gl-exA
go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json
echo

echo "Generated repository full log:"
cd ../gl-exA
git log --oneline --decorate --date-order --reverse --stat --format="%n %h %s%n    Author: %an <%ae> %ad%n    Commit: %cn <%ce> %cd"
echo ""

echo "Generated repository simple log:"
git log --oneline --decorate --date-order --graph
echo ""
