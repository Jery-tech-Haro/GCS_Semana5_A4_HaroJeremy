# CM_STATUS_REGISTER.md — Registro de Estados de Configuración (3.1 Status Accounting)

| EC-ID | Elemento de Configuración | Tipo | Versión/Ref | Estado | Responsable | Evidencia (link/captura) |
|:------|:--------------------------|:-----|:------------|:-------|:------------|:-------------------------|
| **EC-01** | `docs/SRS/SRS_v1.md` | Doc | `v1.1.0` (commit `995b78c`) | Baselined | Analista | commit `995b78c` (ISSUE-1) |
| **EC-02** | `src/app.go` | Code | commit `a87bd13` | Integrado | Dev | commit `a87bd13` |
| **EC-03** | `tests/app_test.go` | Test | commit `a87bd13` | Verificado | QA | `go test ./...` → PASS |
| **EC-04** | `CHANGELOG.md` | Doc | `v1.1.0` (commit `25f9b8d`) | Aprobado | PM | commit `25f9b8d` + release notes |
| **EC-05** | `.gitignore` | Config | commit `1026f5d` | Aprobado | DevOps | commit `1026f5d` (ISSUE-1) |
| **EC-06** | `config/.env.example` | Config | commit `1026f5d` | Integrado | DevOps | commit `1026f5d` (ISSUE-1) |
| **EC-07** | `.github/pull_request_template.md` | Process | commit `308e3c7` | Aprobado | Líder | commit `308e3c7` (ISSUE-1) |
| **EC-08** | `README.md` | Doc | `v1.0.0` (commit `a87bd13`) | Baselined | Equipo | tag `v1.0.0` |
| **EC-09** | `src/app.go` (`FilterByDate`, REQ-003) | Code | `v1.2.0` (rama `audit/req-003-filtro-fecha`) | Verificado | Auditor Funcional | `docs/audit/AUDITORIA_FUNCIONAL_REQ-003.md` + `go test ./...` → PASS |
| **EC-10** | `.github/workflows/ci.yml` | Process | `v1.2.0` (rama `chore/ci-workflow`) | Integrado | Entrega/Despliegue | ejecución de GitHub Actions (ver pestaña Actions) |

*Estados sugeridos:* Registrado · En revisión · Aprobado · Baselined · En implementación · Integrado · Verificado · Liberado · Retirado.

## Notas de Auditoría (ISSUE-1)
- El commit `8705acf` ("actualizar cosas") introdujo `config/.env` con un secreto de ejemplo. Fue corregido en `1026f5d`, que eliminó el archivo del tracking y agregó `config/.env.example` + `.gitignore`.
- El tag `v1.0` (no-SemVer) y `release-1.1` (formato inconsistente) fueron reemplazados por `v1.0.0` y `v1.1.0` respectivamente.
- Todos los commits de corrección referencian ISSUE-1 para trazabilidad.

## Notas de Auditoría (ISSUE-2)
- Auditoría funcional de `REQ-003` (filtro de productos por fecha): se definieron 3 criterios de aceptación y se verificaron con pruebas automatizadas. Detalle en `docs/audit/AUDITORIA_FUNCIONAL_REQ-003.md`.
- `REQ-003` pasa de "pendiente de implementación" a "Implementado y verificado" en `docs/SRS/SRS_v1.md`.

## Notas de Auditoría (ISSUE-3)
- Se agregó pipeline de integración continua (`.github/workflows/ci.yml`) que ejecuta `go build` y `go test` en cada push/PR a `main`, como control de entrega.
