s = input()
mult = {'(': 2, ')': 2, '[': 3, ']': 3}
pair = {'(': ')', '[': ']', ')': '(', ']': '['}
stack = []
ans, cur = 0, 1
for c in s:
    if c in '([':
        stack.append(c)
        cur *= mult[c]
    elif c in ')]':
        if not stack or stack[-1] != pair[c]:
            ans = 0
            break
        if prev == pair[c]:
            ans += cur
        stack.pop()
        cur //= mult[c]
    prev = c

if stack:
    ans = 0

print(ans)
