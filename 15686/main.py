from itertools import combinations

N, M = map(int, input().split())

houses = []
chickens = []
for i in range(N):
    for j, c in enumerate(input().split()):
        if c == '1':
            houses.append((i, j))
        elif c == '2':
            chickens.append((i, j))

best = 2*N*N*N
for mchickens in combinations(chickens, M):
    dist = sum(min(abs(h[0]-c[0]) + abs(h[1]-c[1]) for c in mchickens)
               for h in houses)
    if dist < best:
        best = dist
print(best)
