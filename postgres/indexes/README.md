# Indexes 📇

Hands-on по типам индексов.

```bash
cd postgres/indexes
just up # Postgres на 5433
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

