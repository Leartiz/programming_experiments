# Go scheduler (G / M / P)

## Модель G-M-P

| | |
|---|---|
| **G** (goroutine) | задача: стек, PC, регистры |
| **M** (machine) | OS-thread, на нём реально крутится код |
| **P** (processor) | логический CPU: локальная run queue для G |

**GOMAXPROCS** = число P (по умолчанию = `runtime.NumCPU()`).

```
  P0 [G1 G2 G3]     P1 [G4 G5]
   |                 |
   M (thread)        M (thread)
```

Если у P закончились G - **work stealing** у другого P.

## Goroutine ≠ OS thread

Тысячи goroutine - норм: они дешёвые. M создаются по необходимости (обычно ≈ GOMAXPROCS, пока все заняты).

## Когда goroutine уступает CPU

- `chan` send/receive
- `select`, `time.Sleep`, блокирующий I/O
- sync (mutex, WaitGroup после wake)
- **с Go 1.14+** - асинхронное **preemption**: runtime может прервать долгий CPU-bound цикл (~10 ms)

Старый пример "вечный for без yield" на одном P - сегодня не блокирует навсегда, но **CPU-bound без точек уступа** всё равно плохая идея.

## Happens-before (кратко)

Send на channel → receive на том же channel видит записанные данные. Не замена mutex для произвольного shared state.

## Что смотреть в demo

1. **baseline** - CPU, GOMAXPROCS, goroutines
2. **many sleepers** - много G, мало работы для планировщика
3. **GOMAXPROCS=1** - одна P, G чередуются на time.Sleep
4. **chan rendezvous** - unbuffered chan как точка синхронизации
5. **cpu spin** - один G крутит CPU, другой на sleep (preemption)
