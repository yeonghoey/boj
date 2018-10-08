from collections import defaultdict, deque
N, M, V = map(int, input().split())
g = defaultdict(dict)
for _ in range(M):
    a, b = map(int, input().split())
    g[a][b] = True
    g[b][a] = True

# DFS
dfs, visited, order = deque([V]), set(), []
while len(dfs) > 0:
    v = dfs.pop()
    if v in visited:
        continue
    order.append(v)
    visited.add(v)
    for k in sorted(g[v].keys(), reverse=True):
        if k not in visited:
            dfs.append(k)
print(' '.join(map(str, order)))

# BFS
bfs, visited, order = deque([V]), set(), []
while len(bfs) > 0:
    v = bfs.popleft()
    if v in visited:
        continue
    order.append(v)
    visited.add(v)
    for k in sorted(g[v].keys()):
        if k not in visited:
            bfs.append(k)
print(' '.join(map(str, order)))
