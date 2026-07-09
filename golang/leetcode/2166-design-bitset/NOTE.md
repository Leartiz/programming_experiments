# 2166 Design Bitset

https://leetcode.com/problems/design-bitset/

## Пакеты

- `canonical/` - каноническое решение (два `[]byte` + счётчик).
- `chunkmap/` - lazy `map[chunk]uint64` + `unknownBitVal`.
