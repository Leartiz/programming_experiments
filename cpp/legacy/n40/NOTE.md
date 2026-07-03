## n40 - `QScopedPointer; QPromise; QFuture; QThread`

- implementation [here](./main.cpp)

## Theory

`QScopedPointer` is a small utility class that heavily simplifies this by assigning stack-based memory ownership to heap allocations, more generally called resource acquisition is initialization(RAII).