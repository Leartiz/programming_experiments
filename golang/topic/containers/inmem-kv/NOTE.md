# inmem-kv

In-memory key-value store (stdlib only).

## TODO

- [ ] `map[string]T` + `sync.RWMutex`
- [ ] Get / Set / Delete
- [ ] TTL (optional, см. `play/clean-arch-ws/.../cache/ram`)
- [ ] `v, ok := Get` idiom

## Запуск

```bash
go run .
```

## Потом

HTTP / eviction -> `play/inmem-kv/`
