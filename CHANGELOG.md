# Changelog

## [Unreleased]
- (pendiente)

## [v1.2.0] - 2026-08-24
- Auditoría funcional (ISSUE-2, issue #3, PR #4):
  - Implementado `REQ-003`: filtro de productos por fecha de alta (`FilterByDate`, rango `desde`/`hasta` ISO-8601).
  - Definidos y verificados 3 criterios de aceptación con pruebas automatizadas (`docs/audit/AUDITORIA_FUNCIONAL_REQ-003.md`).
  - `Product` ahora incluye el campo `CreatedAt`.
- Entrega/Despliegue (ISSUE-3, issue #5, PR #6):
  - Agregado pipeline de integración continua (`.github/workflows/ci.yml`) que corre `go build` y `go test ./...` en cada push/PR a `main`.

## [v1.1.0] - 2026-08-21
- Saneamiento y correcciones (ISSUE-1):
  - Eliminado secreto versionado (`config/.env`) del tracking; agregado `config/.env.example` y `.gitignore`.
  - Corregidos tags inconsistentes: `v1.0` → `v1.0.0`, `release-1.1` → `v1.1.0`.
  - Clarificados los criterios pendientes de REQ-003 (filtro de productos por fecha).
  - Completado `CM_STATUS_REGISTER.md` con el registro de estados de los Elementos de Configuración.
  - Agregada plantilla de Pull Request (`.github/pull_request_template.md`).

## [v1.0.0] - 2026-08-21
- Baseline: estructura + SRS v1 + código mínimo (Go) + prueba mínima.
