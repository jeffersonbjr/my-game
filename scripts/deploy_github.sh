#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if ! git remote get-url origin >/dev/null 2>&1; then
  echo "Remote 'origin' não configurado. Configure e faça login no GitHub antes do deploy." >&2
  echo "Exemplo: git remote add origin https://github.com/<usuario>/<repo>.git" >&2
  exit 1
fi

branch="$(git rev-parse --abbrev-ref HEAD)"

./scripts/build_web.sh

echo "Fazendo push da branch '$branch' para origin..."
git push -u origin "$branch"

if command -v gh >/dev/null 2>&1; then
  if gh auth status >/dev/null 2>&1; then
    echo "Disparando workflow de deploy no GitHub Actions..."
    gh workflow run deploy-pages.yml || true
  else
    echo "gh CLI encontrado, mas sem autenticação ativa. Pulei workflow_dispatch."
  fi
else
  echo "gh CLI não encontrado. O deploy ainda acontece automaticamente no push para main/master."
fi

echo "Deploy solicitado. Acompanhe em: $(git remote get-url origin)/actions"
