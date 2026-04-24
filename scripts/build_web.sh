#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DIST_DIR="$ROOT_DIR/web/dist"

rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

GO_WASM_EXEC="$(go env GOROOT)/lib/wasm/wasm_exec.js"
if [[ ! -f "$GO_WASM_EXEC" ]]; then
  GO_WASM_EXEC="$(go env GOROOT)/misc/wasm/wasm_exec.js"
fi

if [[ ! -f "$GO_WASM_EXEC" ]]; then
  echo "wasm_exec.js não encontrado no GOROOT ($(go env GOROOT))." >&2
  exit 1
fi

cp "$ROOT_DIR/web/index.html" "$DIST_DIR/index.html"
cp "$GO_WASM_EXEC" "$DIST_DIR/wasm_exec.js"

GOOS=js GOARCH=wasm go build -o "$DIST_DIR/game.wasm" .

touch "$DIST_DIR/.nojekyll"

echo "Build web concluído em $DIST_DIR"
