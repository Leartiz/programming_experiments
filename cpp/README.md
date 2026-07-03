# C++ ⚙️

## Layout

- `topic/` - language, STL, patterns
- `qt/` - Qt experiments
- `play/` - third-party libraries (boost, spdlog, …)
- `leetcode/` - LeetCode problems
- `legacy/` - old experiments (n*, do not extend)

## New experiment

1. Folder:
   - `leetcode/<num>-<slug>`
   - `qt/<slug>`
   - `play/<slug>`
   - `topic/<theme>/<slug>` when you need grouping
2. `CMakeLists.txt` + `main.cpp` (or `.pro` for Qt)

## Build

```bash
cd topic/random/random-device
cmake -B build
cmake --build build
./build/random-device    # or build/Debug/random-device.exe on Windows
```
