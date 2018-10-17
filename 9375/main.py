from collections import Counter
from itertools import combinations

t = int(input())
while t > 0:
    n = int(input())
    wears = Counter()
    for _ in range(n):
        _, kind = input().split()
        wears[kind] += 1

    x = 1
    for v in wears.values():
        x *= v+1
    print(x-1)

    t -= 1
