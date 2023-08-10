## n16 - problem analysis `Nim Game`

- implementation [here](./main.cpp)

## Explanation

1. The base case: when n = 4, as suggested by the hint from the problem, no matter which number that that first player, the second player would always be able to pick the remaining number.
2. For 1* 4 < n < 2 * 4, (n = 5, 6, 7), the first player can reduce the initial number into 4 accordingly, which will leave the death number 4 to the second player. i.e. The numbers 5, 6, 7 are winning numbers for any player who got it first.
3. Now to the beginning of the next cycle, n = 8, no matter which number that the first player picks, it would always leave the winning numbers (5, 6, 7) to the second player. Therefore, 8 % 4 == 0, again is a death number.
4. Following the second case, for numbers between (2*4 = 8) and (3*4=12), which are 9, 10, 11, are winning numbers for the first player again, because the first player can always reduce the number into the death number 8.

Read more [here](https://leetcode.com/problems/nim-game/solutions/73749/theorem-all-4s-shall-be-false/)