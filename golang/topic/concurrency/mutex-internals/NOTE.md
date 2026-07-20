# Mutex internals

Стек:

```
CPU: CAS / barriers
  -> memory model (когда виден write)
  -> park/unpark (futex / runtime sema)
  -> sync.Mutex
```

## План

1. **CAS без сна** - [../cas-spinlock/](../cas-spinlock/)
2. **memory_order / happens-before** - [../memory-order/](../memory-order/)
3. **toy mutex**: CAS + park - здесь (код позже)
4. **sync.Mutex**: state bits, starvation - NOTE + bench
5. **futex** - [../../../os/TODO.md](../../../os/TODO.md) -> `linux/futex/`

## Идея

- uncontended: один CAS, без syscall
- contended: CAS fail -> sleep в ядре (futex / runtime sema) -> wake
- Lock ≈ acquire, Unlock ≈ release (видимость данных для следующего владельца)
