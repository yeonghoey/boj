MAX_N = 123456*2 + 1
isprime = [True] * (MAX_N)
isprime[0] = False
isprime[1] = False

for i in range(2, MAX_N):
    if not isprime[i]:
        continue
    for j in range(i+i, MAX_N, i):
        isprime[j] = False

while True:
    n = int(input())
    if n == 0:
        break
    ans = sum(1 for i in range(n+1, 2*n+1) if isprime[i])
    print(ans)
