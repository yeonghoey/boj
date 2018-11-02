from random import randint

N = 5000000
L = 1000000
print(f'''{N} {L}
{' '.join([str(randint(-10e9, 10e9)) for i in range(N)])}
''')
