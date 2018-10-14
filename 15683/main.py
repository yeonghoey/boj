from collections import deque
from itertools import product

def D(spec):
    offsets = [(-1, 0), (0, 1), (1, 0), (0, -1)]
    return [
        [
            offsets[i]
            for i, d in enumerate(s)
            if d == '1'
        ] for s in spec
    ]

cctvs = {
    # Top, Right, Down, Left
    '1': D(['1000', '0100', '0010', '0001']),
    '2': D(['1010', '0101']),
    '3': D(['1100', '0110', '0011', '1001']),
    '4': D(['1110', '0111', '1011', '1101']),
    '5': D(['1111']),
}

N, M = [int(s) for s in input().split()]
office = [['x'] + (['0'] * M) + ['x'] if 1 <= i <= N else
          ['x'] * (M*2)
          for i in range(N+2)]

dirs = []
locs = []
total = N*M
for i in range(1, N+1):
    row = input().split()
    for j, ch in enumerate(row, 1):
        office[i][j] = ch
        if ch != '0':
            total -= 1
        if ch in cctvs:
            dirs.append(cctvs[ch])
            locs.append((i, j))

def bfs(start, dirpick, explored):
    q = deque([(start, dir1) for dir1 in dirpick])
    e = set(start)
    x = 0
    while q:
        loc, dir1 = q.popleft()
        loc1 = (loc[0]+dir1[0], loc[1]+dir1[1])
        y1, x1 = loc1
        ch = office[y1][x1]
        if loc1 not in e:
            if ch == '0' or ch in cctvs:
                x += (1 if loc1 not in explored and ch == '0' else 0)
                e.add(loc1)
                q.append((loc1, dir1))
    return x, e

best = 0
for dirpicks in product(*dirs):
    explored = set()
    x = 0
    for start, dirpick in zip(locs, dirpicks):
        x1, e1 = bfs(start, dirpick, explored)
        x += x1
        explored |= e1
    if x > best:
        best = x

print(total - best)
