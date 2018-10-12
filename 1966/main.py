from collections import deque
T = int(input())
for _ in range(T):
    N, M = map(int, input().split())
    A = deque((int(s), i) for i, s in enumerate(input().split()))
    i = -1
    while A:
        maxn, _ = max(A)
        n, i = A.popleft()
        if n >= maxn and i == M:
            break
        if n < maxn:
            A.append((n, i))
    print(N - len(A))
