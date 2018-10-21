def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

N = int(input())
rings = [int(term) for term in input().split()]
first = rings[0]
for ring in rings[1:]:
    g = gcd(first, ring)
    print(f'{first//g}/{ring//g}')
