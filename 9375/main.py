from collections import Counter
from functools import reduce

t = int(input())
while t > 0:
    n = int(input())
    wears = Counter(input().split()[1] for _ in range(n))
    x = reduce(lambda v, e: v*(e+1), wears.values(), 1)
    print(x-1)
    t -= 1
