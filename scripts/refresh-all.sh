#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

OUT="$SCRIPT_DIR/../docs/exA-log.md"

echo "Generating documentation log at $OUT"

echo -n > "$OUT"

echo "# Example rgen run" >> "$OUT"
echo "" >> "$OUT"

cd "$SCRIPT_DIR/../../gl-exA-src"
COMMIT_HASH="$(git rev-parse HEAD)"
COMMIT_HASH_SHORT="$(git rev-parse --short HEAD)"
GITHUB_LINK="https://github.com/mj41/gl-exA-src/tree/$COMMIT_HASH/assets"
echo "Source repository gl-exA-src [commit $COMMIT_HASH_SHORT /assets tree]($GITHUB_LINK) (on GitHub)" >> "$OUT"
echo "" >> "$OUT"

echo "## Script output" >> "$OUT"
echo "\`\`\`" >> "$OUT"
cd "$SCRIPT_DIR"
./refresh-docs-log.sh 2>&1 >> "$OUT"
echo "\`\`\`" >> "$OUT"

echo "## Script code" >> "$OUT"
echo "\`\`\`bash" >> "$OUT"
cat "$SCRIPT_DIR/refresh-docs-log.sh" >> "$OUT"
echo "\`\`\`" >> "$OUT"
echo "" >> "$OUT"

echo "Done. Output written to $OUT"