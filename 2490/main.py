from collections import Counter
for _ in range(3):
    counter = Counter(input().split())
    print('EABCD'[counter['0']])
