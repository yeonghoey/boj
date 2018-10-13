from collections import deque
N, M = [int(s) for s in input().split()]
picks = (int(s) for s in input().split())
q = deque(range(1, N+1))
count = 0
op1 = lambda: q.popleft()
op2 = lambda: q.append(q.popleft())
op3 = lambda: q.appendleft(q.pop())
for p in picks:
    idx = q.index(p)
    pcnt, op = (idx, op2) if idx < len(q)-idx else (len(q)-idx, op3)
    count += pcnt
    while pcnt > 0:
        op()
        pcnt -=1
    op1()
print(count)
