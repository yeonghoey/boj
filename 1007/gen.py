#!/usr/bin/env python
from random import randint

T = 1
N = 20
print(T)
print(N)
for _ in range(N):
    y = randint(-100000, 100000)
    x = randint(-100000, 100000)
    print(f'{y} {x}')
