# Plan de Gestión de Configuración (CM)

## Elementos de Configuración (EC) bajo control
- Documentación (SRS, README, CHANGELOG)
- Código fuente (`src/`)
- Pruebas (`tests/`)
- Configuración (`config/`, `.gitignore`)
- Procesos (plantillas de PR)

## Política de baseline
- Toda baseline formal se marca con un tag SemVer anotado (`vMAJOR.MINOR.PATCH`).
- Ningún EC se considera "Baselined" sin evidencia (commit/PR) y sin registro en `CM_STATUS_REGISTER.md`.

## Trazabilidad
- Todo cambio relevante referencia un Issue (`ISSUE-xx`) en el mensaje de commit y en el Pull Request.
