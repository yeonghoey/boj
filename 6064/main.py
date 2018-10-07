def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

def lcm(a, b):
    return a * b // gcd(a, b)

def solve(m, n, x, y):
    l = lcm(m, n)
    add, base = (m, x) if m > n else (n, y)
    for k in range(base, lcm(m,n)+1, add):
        if (k-1) % m == x-1 and (k-1) % n == y-1:
            return k
    return -1

T = int(input())
for _ in range(T):
    M, N, x, y = map(int, input().split())
    print(solve(M, N, x, y))
