# topic/ - backlog примеров

Очередь идей: дочитал тему -> сделал пример -> перенес в **Сделано** со ссылкой на папку.

Правило: без внешних зависимостей -> `topic/<theme>/<slug>/` (`main.go` + `NOTE.md`).

Play (pgx, prometheus, ...) -> [../play/TODO.md](../play/TODO.md)

OS / kernel (scheduling, /proc, cgroups) -> [../../os/TODO.md](../../os/TODO.md)

## Как пользоваться

1. Записал идею сюда (`[ ]`).
2. Сделал пример -> папка `topic/...`.
3. Убрал из секции выше, добавил одну строку в **Сделано** (общий список, без подразделов).

---

## algorithms / complexity

- [ ] Big O: время + память - шпаргалка и 3-4 типовых кейса
  - план: `topic/algorithms/complexity/`
  - O(n) sliding window vs O(n^2) nested loops
  - extra space: map/slice freq [26] vs sort in-place
  - amortized: append slice, hash map operations
  - связать с leetcode (438, 209) - оценить свой код до/после

## slices / arrays

- [ ] array vs slice: отличия, `==`, copy, pass to func, `[26]int` vs `[]int`
  - план: `topic/slices/array-vs-slice/`
  - value type vs header (ptr+len+cap); array compare `a==b`, slice - loop/`slices.Equal`
  - связь с 438 bench: `[26]int` vs `[]int` freq

- [ ] `append` в функции: slice header копируется, без `return` снаружи не вырастет
  - план: `topic/slices/append-in-function/`
  - bad: `func add(xs []int) { xs = append(xs, 1) }` vs good: `return append(xs, 1)`
  - отдельно: запись `xs[i]=v` видна снаружи (тот же backing array)

- [ ] `append` и cap: in-place vs новый массив - что видно снаружи
  - план: `topic/slices/append-cap-realloc/`
  - cap хватает: пишет в тот же backing array, но len снаружи не меняется (ловушка)
  - cap не хватает: новый array, снаружи старый header - данных не видно
  - demo: `make([]int, 0, 3)` + append в func; сравнить с `xs[:len+1]` / return

## maps

- [ ] внутреннее устройство map: buckets, overflow, hash, рост, iteration order
  - план: `topic/containers/map-internals/`
  - hmap / buckets (8 keys), load factor, evacuate при grow
  - почему ключи comparable; `nil` map vs `make`; concurrent read/write panic
  - ссылки: go blog / runtime map docs; опционально `unsafe` / `reflect` только в NOTE

- [ ] in-memory KV: Get/Set/Delete, map + RWMutex, TTL
  - шаблон: [containers/inmem-kv](./containers/inmem-kv/)

## concurrency / context

- [ ] внутреннее устройство chan: hchan, send/recv, select, close + примитивная реализация
  - план: `topic/concurrency/channel-internals/`
  - buffered vs unbuffered; очередь, park/unpark goroutine
  - минимальный chan: mutex + cond / ring buffer + NOTE со ссылками на runtime

- [ ] внутреннее устройство Mutex: state, sema, starvation mode
  - план: [concurrency/mutex-internals/](./concurrency/mutex-internals/)
  - этап 1: [concurrency/cas-spinlock/](./concurrency/cas-spinlock/)
  - этап 2: [concurrency/memory-order/](./concurrency/memory-order/)
  - fast path / slow path; Mutex vs RWMutex; как связано с runtime sema
  - дальше: toy mutex, futex (os/)

- [ ] atomic ownership: `atomic.Pointer` / refcount vs C++ `shared_ptr`
  - план: `topic/concurrency/atomic-pointer/`
  - Go: GC, зачем `atomic.Pointer[T]`; toy refcount на CAS (сравнение с shared_ptr)
  - связать с cpp: `atomic<shared_ptr>` / refcount + CAS

- [ ] многозадачность: cooperative vs preemptive, Go 1.14+ preemption
  - план: `topic/concurrency/goroutine-scheduling/` (доработать)
  - обновить demo (старый infinite loop на go 1.21+ уже не блокирует)
  - GODEBUG=schedtrace, go tool trace
  - связать с [scheduler/](./concurrency/scheduler/)
  - OS-слой (CFS, /proc): [os/TODO.md](../../os/TODO.md)

- [ ] fan-in: merge нескольких chan в один
  - скелет: [concurrency/merge-channels/](./concurrency/merge-channels/)

- [ ] watcher + context: propagate ctx в фон, graceful close ресурсов
  - план: `topic/concurrency/watcher-context/`

- [ ] значения из отменённого context (`Value`, `Err`) после `cancel()`
  - шаблон: [context/cancelled-values](./context/cancelled-values/)

## graceful_shutdown

- [ ] сравнительная таблица: chan vs WaitGroup vs errgroup vs signal.NotifyContext
  - план: `topic/graceful_shutdown/compare/`
  - есть разрозненные примеры, нет одного NOTE со сводкой

## profiling / runtime

- [ ] размер стека горутины: pprof goroutine profile vs `runtime.Stack` vs `MemStats.StackInuse`
  - план: `topic/concurrency/goroutine-stack/`
  - общий StackSys - не per-goroutine; OS thread ~1-8 MB vs Go dynamic stack

- [ ] bench + pprof walkthrough: `go test -bench` / `go tool pprof` на deep recursion
  - план: `topic/concurrency/goroutine-stack/`

## types / language

- [ ] feature flags: `GOEXPERIMENT` / build tags vs product toggles (env/config)
  - план: `topic/language/feature-flags/`

## architecture / patterns

- [ ] connascence: algorithm, identity (dynamic) - ещё примеры в static-connascence

- [ ] слоистая архитектура: handler -> use case -> repo - где граница tx
  - план: `topic/architecture/layered-tx-boundary/`
  - tx открывает use case, repo принимает `Querier`/`Tx`, не handler

- [ ] паттерны: Repository, Unit of Work, WithTx - один NOTE + минимальный код
  - план: `topic/architecture/patterns-repo-uow/`

- [ ] handler / service / repository - кто за что отвечает (чеклист на собес)
  - план: `topic/architecture/layered-roles/`

## Сделано

- [x] defer cancel / context shutdown - [concurrency/context-shutdown/](./concurrency/context-shutdown/)
- [x] loop + closure в for - [concurrency/loop-closure/](./concurrency/loop-closure/)
- [x] sizeof struct{} vs any - [types/empty-size/](./types/empty-size/)
- [x] static connascence - [architecture/static-connascence/](./architecture/static-connascence/)
- [x] graceful shutdown variants - [graceful_shutdown/](./graceful_shutdown/)
- [x] G/M/P scheduler - [concurrency/scheduler/](./concurrency/scheduler/)
- [x] priority queue / heap - [containers/priority-queue/](./containers/priority-queue/)
- [x] rune vs UTF-8 in string - [language/rune-utf8/](./language/rune-utf8/)
- [x] slice len=0 через reslice - [slices/slice-len-zero-cap/](./slices/slice-len-zero-cap/)
