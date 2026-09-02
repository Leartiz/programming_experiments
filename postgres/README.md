# postgres/

Hands-on SQL (PostgreSQL). Чистый SQL здесь; pgx/docker в [golang/play/](../golang/play/).

Backlog: [TODO.md](./TODO.md)

## Старт

```bash
cd postgres/indexes
just up
just psql
```

## Уже есть

- [indexes/](./indexes/) — индексы, EXPLAIN
- [isolation/](./isolation/) — зачаток tx isolation
- [ddl/n1/](./ddl/n1/) — up/down SQL
