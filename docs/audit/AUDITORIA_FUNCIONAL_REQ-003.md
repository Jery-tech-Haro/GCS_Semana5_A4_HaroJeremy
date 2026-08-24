# Auditoría Funcional — REQ-003 (filtro de productos por fecha)

Rol: Auditor Funcional (Requisitos)
Requisito auditado: `REQ-003` ([docs/SRS/SRS_v1.md](../SRS/SRS_v1.md))
Issue: `ISSUE-2` (GitHub #3) — PR #4

## Criterios de aceptación

| # | Criterio | Resultado |
|:--|:---------|:----------|
| 1 | Un rango `desde`/`hasta` en formato ISO-8601 (`YYYY-MM-DD`) retorna solo los productos con fecha de alta dentro del rango. | ✅ PASA (`TestFilterByDateRange`) |
| 2 | Sin criterio de fecha (`desde` y `hasta` vacíos), se retorna el listado completo de productos. | ✅ PASA (`TestFilterByDateEmptyReturnsAll`) |
| 3 | Una fecha inválida o `desde > hasta` retorna un error controlado (no un panic ni un resultado silencioso incorrecto). | ✅ PASA (`TestFilterByDateInvalidRange`, `TestAddProductInvalidDate`) |

## Evidencia de ejecución

```
$ go test ./... -v
=== RUN   TestAddAndList
--- PASS: TestAddAndList (0.00s)
=== RUN   TestAddProductInvalidDate
--- PASS: TestAddProductInvalidDate (0.00s)
=== RUN   TestFilterByDateRange
--- PASS: TestFilterByDateRange (0.00s)
=== RUN   TestFilterByDateEmptyReturnsAll
--- PASS: TestFilterByDateEmptyReturnsAll (0.00s)
=== RUN   TestFilterByDateInvalidRange
--- PASS: TestFilterByDateInvalidRange (0.00s)
PASS
ok  	github.com/Jery-tech-Haro/GCS_Semana5_A4_HaroJeremy/tests	0.293s
```

## Conclusión

`REQ-003` cumple los 3 criterios de aceptación definidos. Se implementó `FilterByDate(desde, hasta string) ([]Product, error)` en [src/app.go](../../src/app.go), y `Product` ahora incluye `CreatedAt` (ISO-8601). El requisito pasa de estado "pendiente de implementación" a "Implementado y verificado" en el SRS y en `CM_STATUS_REGISTER.md`.
