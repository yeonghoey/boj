a = int(input())
b = int(input())
c = int(input())
p = a*b*c

counts = [0] * 10
for c in str(p):
    counts[int(c)] += 1

for n in counts:
    print(n)
