# asio::deferred

Completion token (Boost.Asio, C++20): `async_*` не стартует сразу, а возвращает
отложенную операцию. Её можно запустить позже - handler'ом или через `co_await`.

Нужны Boost.Asio и C++20.

```bash
cd cpp/play/asio-deferred
cmake -B build -build
./build/asio-deferred
```
