# std::forward

Имя параметра внутри функции - всегда lvalue. `std::forward<T>(x)` возвращает
ту же value category, с которой аргумент пришёл в шаблон.

1. `&` / `&&` у `operator()` - какая перегрузка вызовется.
2. `Heavy` - rvalue с `forward` даёт **move**, без него - **copy** (видно по счётчикам и ms).

```bash
cd cpp/topic/language/forward
cmake -B build && cmake --build build
./build/forward
```
