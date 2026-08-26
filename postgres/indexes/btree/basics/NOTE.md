# B-tree: basics

Seq Scan vs Index Scan, `EXPLAIN (ANALYZE)`, когда индекс не помогает.

## Запуск

```bash
cd postgres/indexes/btree/basics
just demo          # up + setup + explain
# или по шагам:
just setup
just explain       # идемпотентно
```

Postgres уже поднят?

Достаточно `just setup` / `just explain` (контейнер из `indexes/`, порт 5433).

## Что смотреть в плане

| шаг | ожидание |
|-----|----------|
| `WHERE email = ...` без индекса | `Seq Scan` |
| после `CREATE INDEX ... ON (email)` | `Index Scan` |
| `WHERE lower(email) = ...` | индекс по `email` не подходит |
| `WHERE email LIKE '%...'` | leading `%` -> Seq Scan |
| `WHERE email LIKE 'user5000%'` | на 10k часто Seq Scan; см. `btree/pattern-ops/` |
