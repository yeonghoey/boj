MAX_N = 10000 + 1
isprime = [True] * (MAX_N)
isprime[0] = False
isprime[1] = False

for i in range(2, MAX_N):
    if not isprime[i]:
        continue
    for j in range(i+i, MAX_N, i):
        isprime[j] = False

T = int(input())
for _ in range(T):
    n = int(input())
    for i in range(n//2, 1, -1):
        if isprime[i] and isprime[n-i]:
            print('%d %d' % (i, n-i))
            break
