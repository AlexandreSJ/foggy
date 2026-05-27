# Exercise 04

[Go Back](README.md)

## Description

> This problem was found here: https://codeforces.com/problemset/problem/2167/A

You are given 4 sticks of lengths a, b, c, and d. You can not break or bend them.

Determine whether it is possible to form a square∗ using the given sticks.


## Testing Input

The first line contains a single integer t (1≤t≤10^4) — the number of test cases.

The only line of each test case contains four integers a, b, c, and d (1≤a,b,c,d≤10) — the lengths of the sticks.

Given this input, the user should expect the following [Expected Output](#expected-output):

```bash
7
1 2 3 4
1 1 1 1
2 2 2 2
1 2 1 2
1 1 5 5
5 5 5 5
4 10 5 9
```

## Expected Output

For each test case, print "YES" if it is possible to form a square using the given sticks, and "NO" otherwise.

You may print each letter in any case (uppercase or lowercase). For example, the strings "yEs", "yes", "Yes", and "YES" will all be recognized as a positive answer.

```bash
NO
YES
YES
NO
NO
YES
NO
```