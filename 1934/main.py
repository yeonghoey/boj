def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

def lcm(a, b):
    return a * b // gcd(a, b)

T = int(input())
for _ in range(T):
    A, B = (int(term) for term in input().split())
    print(lcm(A, B))
