from collections import deque
T = int(input())
while T > 0:
    p = input()
    n = int(input())
    q = deque(s for s in input()[1:-1].split(',') if s)
    rev = False
    err = False
    for op in p:
        if op == 'R':
            rev = not rev
        elif op == 'D':
            if not q:
                err = True
                break
            if rev:
                q.pop()
            else:
                q.popleft()
        else:
            pass

    result = list(q if not rev else reversed(q))
    ans = (f'[{",".join(result)}]' if not err else 'error')
    print(ans)
    T -= 1
