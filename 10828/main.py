N = int(input())
stack = []
for _ in range(N):
    line = input().split()
    cmd = line[0]
    if cmd == 'push':
        arg = int(line[1])
        stack.append(arg)
    elif cmd == 'pop':
        print(stack.pop() if stack else -1)
    elif cmd == 'size':
        print(len(stack))
    elif cmd == 'empty':
        print(0 if stack else 1)
    elif cmd == 'top':
        print(stack[-1] if stack else -1)
    else:
        pass
