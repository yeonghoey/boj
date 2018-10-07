from collections import Counter

N = int(input())
A = Counter()

for _ in range(N):
    a = int(input())
    A[a] += 1

for n, count in sorted(A.items()):
    for _ in range(count):
        print(n)
