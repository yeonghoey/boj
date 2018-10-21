def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

def lcm(a, b, g):
    return a * b // g

A, B = (int(term) for term in input().split())
g = gcd(A, B)
print(g)
print(lcm(A, B, g))
