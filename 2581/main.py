M = int(input())
N = int(input())
isprime = [True] * (N+1)
isprime[0] = False
isprime[1] = False

pmin, psum = -1, 0
for i in range(2, N+1):
    if not isprime[i]:
        continue
    if i >= M:
        pmin = i if pmin == -1 else pmin
        psum += i
    for j in range(i+i, N+1, i):
        isprime[j] = False

if psum > 0:
    print(psum)
print(pmin)
