n = int(input())
stack, last, out = [0], 1, []
for _ in range(n):
    a = int(input())
    while last <= a:
        stack.append(last)
        last += 1
        out.append('+')
    if stack[-1] == a:
        stack.pop()
        out.append('-')
    else:
        out = ['NO']
        break
print('\n'.join(out))
