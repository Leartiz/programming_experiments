# memory_order

## 1. relaxed — counter

`N` threads: `fetch_add(1, memory_order_relaxed)`.
Result is always `N * iterations` — atomicity of the value, no sync with other cells.

Use for hit counters / metrics where only the number matters.

## 2. release / acquire — publish payload

```
writer                         reader
──────                         ──────
data = 42;                     while (!ready.load(acquire))
ready.store(true, release);        ;
                               use(data);  // guaranteed 42
```

- **release** store: publishes all prior writes in this thread to a synchronizing acquire.
- **acquire** load: if it sees that store, those prior writes are visible.

`data` is a plain `int`. The flag is the sync point.

### Why not relaxed on the flag

```cpp
data = 42;
ready.store(true, std::memory_order_relaxed);  // no publish

if (ready.load(std::memory_order_relaxed))
    use(data);  // may not be 42 (formally)
```

On x86 this often “works by accident”; the model still forbids relying on it.

## 3. seq_cst — one global order

Default if you omit the order argument.

Release/acquire syncs **along a chain** (writer → reader on the same flag).
It does **not** put all atomic ops from all threads into one total order.

Classic store buffering:

```
thread A              thread B
x = true;             y = true;
r1 = y;               r2 = x;
```

Question: can `r1 == false && r2 == false`?

| order | answer |
|-------|--------|
| **seq_cst** on all four ops | **no** — forbidden |
| **release** store + **acquire** load | **yes** — allowed by the model |

Demo counts that outcome over many runs.

**On x86** both counts are often `0` (TSO is strong). That does not make release/acquire “as strong as seq_cst” — only that the machine rarely shows the allowed reorder. On ARM the second count can be `> 0`.

When you need “everyone agrees on one timeline of these atomics” → `seq_cst`.
For publish-a-payload, release/acquire is enough and usually cheaper.

## Run

```bash
cmake -B build && cmake --build build
./build/Debug/memory-order.exe   # Windows MSVC
# or ./build/memory-order
```

Go parallel (Mutex ≈ acquire/release):
[../../../golang/topic/concurrency/memory-order/](../../../golang/topic/concurrency/memory-order/)
