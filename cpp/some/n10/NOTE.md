## n10 - problem analysis `Contains Duplicate II`

- implementation [here](./main.cpp)

## Visualization

### Examples (with errors?)

```
using value = int
using index = int
```

#### №1

```
mp<value, index> = {}
input: [1 2 3 1], k = 3 <-------------------------
                                                 |
loop:                                            |
    #0 (end iteration):                          |
        mp = { 1: 0 }                            |
    #1:                                          |
        mp = { 1: 0, 2: 1 }                      |
    #2:                                          |
        mp = { 1: 0, 2: 1, 3: 2 }                |
    #3:                                          |
        mp = { |1|: 0, 2: 1, 3: 2, |1|: 3 }      |
        k == 3 - 0 == 3 then ---                 |
                               |                 |
        return true <-----------------------------

```

#### №2

```
mp<value, index> = {}
input: [1 2 3 1 2 3], k = 2 <----------------------
                                                  |
loop:                                             |
    #0 (end iteration):                           |
        mp = { 1: 0 }                             |
    #1:                                           |
        mp = { 1: 0, 2: 1 }                       |
    #2:                                           |
        mp = { 1: 0, 2: 1, 3: 2 }                 |
    #3:                                           |
        mp = { |1|: 0, 2: 1, 3: 2, |1|: 3 }       |
        k == 3 - 0 == 3 != 2 ------------------------------
    #4:                                                   |
        mp = { 1: 0, |2|: 1, 3: 2, 1: 3, |2|: 4 }         | 
        k == 4 - 1 == 3 != 2 ------------------------------
    #5:                                                   |
        mp = { 1: 0, 2: 1, |3|: 2, 1: 3, 2: 4, |3|: 5 }   | 
        k == 5 - 2 == 3 != 2 ------------------------------

return false

```