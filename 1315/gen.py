from random import randint
N = 100
quests = '\n'.join([f'{randint(1, 100)} {randint(1, 100)} {randint(1, 2)}' for i in range(1, N+1)])
print(f'''{N}
{quests}''')
