from random import randint
T = 100
print(T)
for _ in range(T):
    n = randint(1, 100000)
    print(n)
    print(' '.join(str(randint(1, 100)) for _ in range(n)))
    print(' '.join(str(randint(1, 100)) for _ in range(n)))
