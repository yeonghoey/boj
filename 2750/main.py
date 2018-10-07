N = int(input())
a = [int(input()) for _ in range(N)]

for i in range(1, N):
    for j in reversed(range(1, i+1)):
        if a[j-1] > a[j]:
            a[j-1], a[j] = a[j], a[j-1]
        else:
            break
for n in a:
    print(n)
