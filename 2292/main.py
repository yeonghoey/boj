x = int(input())
cur, d = 1, 6
count = 1
while x > cur:
    cur += d
    d += 6
    count += 1
print(count)
