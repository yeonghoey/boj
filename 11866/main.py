from collections import deque
N, M = map(int, input().split())
i, o, q = 0, [], deque(range(1, N+1))
while q:
    i += 1
    a = q.popleft()
    if i % M == 0:
        o.append(str(a))
    else:
        q.append(a)
print(f'<{", ".join(o)}>')
