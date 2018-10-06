t = int(input())
for _ in range(t):
    r, s = input().split()
    r = int(r)
    print(''.join(c for c in s for _ in range(r)))
