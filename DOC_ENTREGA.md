# DOC_ENTREGA.md — Entrega/Despliegue

Rol: Entrega/Despliegue (Opcional PRO)
Issue: `ISSUE-3` (GitHub #5) — PR #6

## Pipeline de integración continua

`.github/workflows/ci.yml` ejecuta automáticamente en cada `push` y `pull_request` hacia `main`:
1. `go build ./...`
2. `go test ./... -v`

Evidencia de ejecución: pestaña **Actions** del repositorio en GitHub, tras el push de esta rama y la apertura del PR.

## Pasos de entrega manual (equivalentes, sin CI)

Si se necesita validar localmente antes de un release:

```bash
go build ./...
go test ./... -v
```

Ambos comandos deben terminar sin errores (`PASS` en todos los tests) antes de aprobar un merge a `main` o de crear un tag de release.

## Checklist de entrega

- [ ] `go build ./...` sin errores.
- [ ] `go test ./... -v` con todos los tests en `PASS`.
- [ ] Workflow de GitHub Actions en verde (ver pestaña Actions).
- [ ] `CHANGELOG.md` actualizado con la versión correspondiente.
