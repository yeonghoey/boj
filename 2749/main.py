from functools import lru_cache

base = [1, 1, 1, 0]

@lru_cache()
def fib(n):
    if n == 0:
        return [1, 0, 0, 1]
    if n == 1:
        return base

    half = fib(n // 2)
    half2 = mul(half, half)
    if n % 2 == 0:
        return half2
    else:
        return mul(half2, base)

def mul(mx, my):
    xa, xb, xc, xd = mx
    ya, yb, yc, yd = my
    return [
        (xa*ya + xb*yc) % 1000000,
        (xa*yb + xb*yd) % 1000000,
        (xc*ya + xd*yc) % 1000000,
        (xc*yb + xd*yd) % 1000000,
    ]

n = int(input())
print(fib(n)[1])
