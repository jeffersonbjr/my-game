# my-game

Jogo simples feito com Ebiten.

## Rodar localmente (desktop)

```bash
go run .
```

## Rodar na web (WebAssembly)

### Build local

```bash
./scripts/build_web.sh
```

Depois sirva `web/dist` com um servidor HTTP:

```bash
cd web/dist
python3 -m http.server 8000
```

Abra `http://localhost:8000`.

## Publicar no GitHub Pages

1. Crie o repositório no GitHub e faça push do código.
2. Em **Settings → Pages**, escolha **GitHub Actions** como source.
3. Faça push para `main` ou `master` (ou rode manualmente o workflow).
4. Aguarde o workflow `Deploy web build to GitHub Pages` concluir.
5. A URL final aparecerá no job de deploy (campo `page_url`).

### Deploy com um comando

```bash
./scripts/deploy_github.sh
```

O script:
- gera o build web;
- faz push da branch atual;
- tenta disparar o workflow `deploy-pages.yml` via `gh` (se estiver autenticado).

### Comandos úteis (primeiro push)

```bash
git remote add origin https://github.com/<seu-usuario>/<seu-repo>.git
git branch -M main
git push -u origin main
```
