from collections import Counter
T = int(input())
for _ in range(T):
    s = input()
    l = 0
    for c in s:
        l = l+1 if c == '(' else l-1
        if l < 0:
            break
    print('YES' if l == 0 else 'NO')
