N, K = (int(s) for s in input().split())
x, y = 1, 1
for k in range(1, K+1):
    x *= (N-k+1)
    y *= k
print(x // y)
