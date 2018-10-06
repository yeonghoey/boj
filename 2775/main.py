from itertools import accumulate

T = int(input())
for _ in range(T):
    k = int(input())
    n = int(input())
    a = range(1, n+1)
    for _ in range(1, k):
        a = accumulate(a)
    print(sum(a))
