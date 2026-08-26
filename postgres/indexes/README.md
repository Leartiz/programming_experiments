# postgres/indexes

Hands-on по типам индексов. Postgres на **5433**, docker + just.

```bash
cd postgres/indexes
just up
```

## Layout

| | путь | тема |
|---|------|------|
| ✅ | [btree/basics/](./btree/basics/) | Seq Scan vs Index Scan, EXPLAIN |
| ☐ | [btree/composite/](./btree/composite/) | `(a, b)`, leftmost prefix |
| ☐ | [btree/partial/](./btree/partial/) | partial index |
| ☐ | [btree/covering/](./btree/covering/) | `INCLUDE`, Index Only Scan |
| ☐ | [btree/pattern-ops/](./btree/pattern-ops/) | `text_pattern_ops`, `LIKE` |
| ☐ | [gin/jsonb/](./gin/jsonb/) | GIN на jsonb |
| ☐ | [gist/](./gist/) | GiST (geo, ranges) |

## Just

| где | команды |
|-----|---------|
| `indexes/` | `up`, `down`, `psql` |
| `btree/basics/` | `setup`, `explain`, `demo` |

