from collections import deque
N = int(input())
q = deque()
for _ in range(N):
    terms = input().split()
    op = terms[0]
    if op == 'push_front':
        q.appendleft(terms[1])
    elif op == 'push_back':
        q.append(terms[1])
    elif op == 'pop_front':
        print(q.popleft() if q else -1)
    elif op == 'pop_back':
        print(q.pop() if q else -1)
    elif op == 'size':
        print(len(q))
    elif op == 'empty':
        print(1 if not q else 0)
    elif op == 'front':
        print(q[0] if q else -1)
    elif op == 'back':
        print(q[-1] if q else -1)
    else:
        pass
