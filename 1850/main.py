def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

A, B = (int(term) for term in input().split())
print(''.join('1' for _ in range(gcd(A, B))))
