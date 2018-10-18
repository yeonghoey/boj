from collections import deque
n = int(input())
a = deque([(0, 0), (0, 0)], 2)
for _ in range(n):
    k = int(input())
    a.append((max(a[-2])+k, a[-1][0]+k))
print(max(a[-1]))
