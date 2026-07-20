# memory-order / happens-before

Mutex - не только "один в секции", но и **видимость данных**.

- Unlock ≈ **release** - опубликовал записи до unlock
- Lock ≈ **acquire** - увидел всё, что опубликовал предыдущий Unlock

Без sync компилятор/CPU не обязаны показывать reader'у записи writer'а в ожидаемом виде -> data race.

В Go имена `memory_order_*` в API нет; у `sync.Mutex` семантика acquire/release внутри.
В C++ их видно явно в `std::atomic` (`std::memory_order_acquire` / `release`).

## Demo

```bash
go run -race .              # с Mutex - чисто
go run -race . -mutex=false # data race в отчёте
```

`-race` - основной способ "увидеть глазами" на Go.
