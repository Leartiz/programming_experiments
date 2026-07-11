#pragma once

// NOTE: C(n, k) = n! / (k! * (n-k)!)
//       how many ways to pick k items from n (order ignored).
//
//   Loop (i = 0 .. k-1):  res = res * (n-i) / (i+1)
//
//   Example C(5, 2) — pick 2 from five items {A,B,C,D,E}:
//
//     pool:   A  B  C  D  E        n=5, k=2
//             *  *                 one pair of slots
//
//     i=0:  res = 1 * 5/1 = 5     (C(5,1): pick 1st item, 5 choices)
//     i=1:  res = 5 * 4/2 = 10    (C(5,2): add 2nd, divide order k!)
//
//     all pairs (10):
//       AB AC AD AE | BC BD BE | CD CE | DE
//

// NOTE:
//
//   факториал нуля всегда равен единице - 0! = 1
//   факториал единицы всегда равен единице - 1! = 1
//
//   n! = 1 * 2 * 3 * n
//   k! = 1 * 2 * 3 * k
//
//   C(n, k) = n! / (k! * (n-k)!)

// NOTE:
/*
            n * (n-1) * ... * (n-k+1)
    C(n,k) = -------------------------
                k!

    C(5,2) = 5 * 4 * 3 * 2 * 1     5 * 4     <- только 2 верхних
            -----------------  =  -----  =  10
            (2*1) * (3*2*1)       2*1
              k!   (n-k)!
*/


// NOTE:
/*
    5, 2
    n - k = 5 - 2 = 3

    i < 2 {
        res = 1 * 5/1 = 5
        res = 5 * 4/2 = 10
    }

    что-то осталось всё таки не понятно!
*/
long long combinations(int n, int k)
{
    if (k < 0 || k > n)
        return 0;

    // 5 - 2 = 3
    // 5 - 4 = 1

    // C(n, k) = C(n, n-k)
    if (k > n - k)
        k = n - k;

    long long res = 1;
    for (int i = 0; i < k; ++i)
    {
        res = res * (n - i) / (i + 1);
    }
    return res;
}
