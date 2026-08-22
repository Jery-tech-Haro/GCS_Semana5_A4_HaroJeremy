# Changelog

## [Unreleased]
- (pendiente)

## [v1.1.0] - 2026-08-21
- Saneamiento y correcciones (ISSUE-1):
  - Eliminado secreto versionado (`config/.env`) del tracking; agregado `config/.env.example` y `.gitignore`.
  - Corregidos tags inconsistentes: `v1.0` → `v1.0.0`, `release-1.1` → `v1.1.0`.
  - Clarificados los criterios pendientes de REQ-003 (filtro de productos por fecha).
  - Completado `CM_STATUS_REGISTER.md` con el registro de estados de los Elementos de Configuración.
  - Agregada plantilla de Pull Request (`.github/pull_request_template.md`).

## [v1.0.0] - 2026-08-21
- Baseline: estructura + SRS v1 + código mínimo (Go) + prueba mínima.
