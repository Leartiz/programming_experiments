## n37 - `QPromise; QFuture; QFutureWatcher; QtConcurrent`

- implementation [here](./main.cpp)

## Theory

### `QFuture`

`QFuture` class represents the result of an asynchronous computation.
`QFuture` allows threads to be synchronized against one or more results which will be ready at a later point in time.
If a result is not available at the time of calling the `result()`, `resultAt()`, `results()` and `takeResult()` functions, `QFuture` *will wait until the result becomes available*.
`QFuture` provides a convenient (удобный) way of chaining (связывание) multiple sequential computations using `then()`.
To interact with running tasks using signals and slots, use `QFutureWatcher`.

### `QPromise`

`QPromise` class provides a way to store computation results to be accessed by `QFuture`.
`QPromise` provides a simple way to communicate progress and results of the user-defined computation to `QFuture` in an asynchronous fashion. For the communication to work, `QFuture` must be constructed by `QPromise`.
By design, `QPromise` is a move-only object.