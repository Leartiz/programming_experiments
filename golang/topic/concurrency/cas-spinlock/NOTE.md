# CAS spinlock

CAS = Compare-And-Swap: атомарно "если было expected - записать new".
CAS = Compare-And-Swap = "сравни и при успехе замени".

- Lock: пока `CAS(0 -> 1)` не успешен - крутимся (spin).
- Unlock: `store 0`.

Плюс: просто, без ядра ОС.
Минус: waiter жрёт CPU.

Связь с mutex: fast path тот же (CAS), вместо spin - park.

## ASM (схема, не полный листинг)

| arch | идея |
|---|---|
| amd64 | `lock cmpxchg` |
| arm64 | `ldxr` + `stxr` (цикл) |

## Race detector

```bash
go run -race .
```

Проверяет data race на shared memory. Молчит → доступы к `n` под lock согласованы. Это не формальное доказательство, но основной практический инструмент в Go.

## Bench

```bash
go test "-bench=." "-cpu=1,4"
```

Сравнивает spinlock и `sync.Mutex`. На `-cpu=4` при contention spin обычно хуже (busy-wait жрёт CPU).
