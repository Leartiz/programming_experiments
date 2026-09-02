# postgres/ - SQL backlog

Очередь: дочитал тему -> SQL demo в `postgres/...` -> Go в [golang/play/](../golang/play/TODO.md) при необходимости.

Docker: [indexes/docker-compose.yml](./indexes/docker-compose.yml) (порт **5433**), `cd postgres/indexes && just up`.

## Как пользоваться

1. Идея сюда (`[ ]`).
2. Чистый SQL -> `postgres/<theme>/<slug>/` (`README.md` + `*.sql`, опционально `Justfile`).
3. Приложение (pgx, две goroutine) -> `golang/play/postgres-tx/...`.
4. Готово -> строка в **Сделано**, убрать из очереди.

---

## 1. transactions (начать здесь)

- [ ] isolation: READ COMMITTED vs REPEATABLE READ (две сессии)
  - план: `isolation/rc-vs-rr/` — зачаток: [isolation/](./isolation/)
  - phantom / non-repeatable read на одном сценарии
  - Go: [play/postgres-tx/isolation-levels](../golang/play/TODO.md)

- [ ] `SELECT ... FOR UPDATE` vs без lock (lost update)
  - план: `transactions/for-update/`
  - связать с [play/select-two-updates](../golang/play/postgres-tx/select-two-updates/)
  - Go: `play/postgres-tx/select-for-update/`

- [ ] deadlock — два UPDATE в разном порядке
  - план: `transactions/deadlock/`
  - `pg_locks`, как Postgres выбирает victim

## 2. indexes / EXPLAIN

- [x] B-tree basics: Seq Scan vs Index Scan
  - [indexes/btree/basics/](./indexes/btree/basics/) — `just demo`

- [ ] composite index + leftmost prefix
  - план: [indexes/btree/composite/](./indexes/btree/composite/)

- [ ] partial index (`WHERE` в определении)
  - план: [indexes/btree/partial/](./indexes/btree/partial/)

- [ ] covering / `INCLUDE` → Index Only Scan
  - план: [indexes/btree/covering/](./indexes/btree/covering/)

- [ ] `text_pattern_ops`, `LIKE 'prefix%'`
  - план: [indexes/btree/pattern-ops/](./indexes/btree/pattern-ops/)

- [ ] JOIN + план (nested loop / hash / merge)
  - план: `queries/join-explain/`
  - два таблицы, `EXPLAIN (ANALYZE)` на разных JOIN

- [ ] GIN на jsonb (`@>`, `?`)
  - план: [indexes/gin/jsonb/](./indexes/gin/jsonb/)

## 3. queries (база SQL)

- [ ] JOIN шпаргалка: INNER / LEFT, когда что
  - план: `queries/joins-basics/`
  - маленькая схема orders + users

- [ ] GROUP BY + HAVING + агрегаты
  - план: `queries/group-by/`

- [ ] подзапросы vs `EXISTS` vs JOIN
  - план: `queries/exists-vs-join/`

- [ ] `NULL`, `UNIQUE`, partial unique (ловушки)
  - план: `queries/null-unique/`

## 4. design / DDL

- [ ] нормализация vs денормализация — один пример
  - план: `ddl/normalization-sketch/`
  - есть зачаток: [ddl/n1/](./ddl/n1/)

- [ ] миграции up/down (минимальный паттерн)
  - расширить [ddl/n1/](./ddl/n1/)

## 5. app layer (не чистый SQL)

- [ ] layered tx: handler -> use case -> repo на `pgx.Tx`
  - [play/postgres-tx/layered-with-tx/](../golang/play/TODO.md)
  - topic: architecture/layered-tx-boundary

- [ ] connection pool: лимиты, `too many clients`
  - уже: [play/pgx-pool](../golang/play/pgx-pool/)

## Сделано

- [x] B-tree + EXPLAIN basics — [indexes/btree/basics/](./indexes/btree/basics/)
- [x] два UPDATE, две tx (sketch) — [golang/play/postgres-tx/select-two-updates](../golang/play/postgres-tx/select-two-updates/)
