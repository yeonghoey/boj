M, N = map(int, input().split())
isprime = [True] * (N+1)
isprime[0] = False
isprime[1] = False

for i in range(2, N+1):
    if not isprime[i]:
        continue
    if i >= M:
        print(i)
    for j in range(i+i, N+1, i):
        isprime[j] = False
