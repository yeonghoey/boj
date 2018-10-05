def judge(a, b):
    return ('ascending' if a < b else 'descending')

a = input().split()
it = iter(a)
first = next(it)
prev = next(it)
ans = judge(first, prev)
for c in it:
    current = judge(prev, c)
    if current != ans:
        ans = 'mixed'
        break
    prev = c
print(ans)
