# API Inventario (mini)
- Endpoints simulados: GET /products, POST /products (lógica implementada en `src/app.go`)
- Objetivo: repo auditable (versiones + estados + trazabilidad)

## Cómo ejecutar
No se requiere despliegue real. La lógica de negocio está implementada en Go y se valida con pruebas unitarias:

```bash
go test ./...
```

## Convención
- Commits: chore/docs/feat/fix + referencia ISSUE-xx
- Versiones: SemVer (vMAJOR.MINOR.PATCH)
