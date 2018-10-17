N = int(input())

x = 1
for n in range(1, N+1):
    x *= n

count = 0
for ch in reversed(str(x)):
    if ch != '0':
        break
    else:
        count += 1
print(count)
