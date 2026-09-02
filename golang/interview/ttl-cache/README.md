# ttl-cache

In-memory кэш с TTL и фоновой очисткой.

## Запуск

```bash
cd golang/interview/ttl-cache
go run .
```

## Задача

- найти проблемы (конкурентность, lifecycle, ресурсы)
- улучшить API (`interface{}`, shutdown, тестируемость)
- опционально: generics, `RWMutex`, бенч на hot path

Оригинал: [cache/cache.go](./cache/cache.go).
