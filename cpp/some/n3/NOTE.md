## n3 - problem analysis `Combination Sum`

- implementation [here](./main.cpp)

## Visualization

### Example (with errors?)
```
[1, 2, 3] -> 2

    |
    V

candidates: [1, 2, 3]
target: 2

#1 - recursion:
    result: []
    current: []
    interimSum: 0
    index: 0

    ---------------------------------------------

    #1 - loop:
        one: 1
        next: [1]
        lt. target -> 

        #2 - recursion:
            result: []
            current: [1]
            interimSum: 1
            index: 0

                #1 - loop:
                    one: 1
                    next: [1 1]
                    result: [[1 1]] 
                    eq. target -> OK

                #2 - loop:
                    one: 2
                    next: [1 2]
                    result: [[1 1]]
                    gt. target -> FAIL

                #3 - loop:
                    one: 3
                    next: [1 3]
                    result: [[1 1]]
                    gt. target -> FAIL

    ---------------------------------------------

    #2 - loop:
        one: 2
        next: [2]
        result: [[1 1] [2]] 
        eq. target -> OK

    #3 - loop:
        one: 3
        next: [3]
        result: [[1 1] [2]] 
        gt. target -> FAIL

```