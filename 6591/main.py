def comb(n, k):
    k = min(k, n-k)
    kl = list(range(2, k+1))
    ki = 0
    x = 1
    for ni in range(n, n-k, -1):
        x *= ni
        while ki < len(kl) and x % kl[ki] == 0:
            x //= kl[ki]
            ki += 1
    return x

while True:
    n, k = (int(s) for s in input().split())
    if n == 0 and k == 0:
        break
    print(comb(n, k))
