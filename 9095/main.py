T = int(input())
a = [0, 1, 2, 4]
for i in range(4, 11):
    a.append(a[i-3] + a[i-2] + a[i-1])
for _ in range(T):
    n = int(input())
    print(a[n])
