from collections import deque
N, K = map(int, input().split())
q = deque([(N, 0)])
v = set([N])
best = abs(K-N)
while q:
    n, d = q.popleft()
    if n == K:
        best = d
        break
    if d > best:
        break
    for nn in [n-1, n+1, n*2]:
        if nn >= 0 and nn <= K+N and nn not in v:
            v.add(nn)
            q.append((nn, d+1))
print(best)
