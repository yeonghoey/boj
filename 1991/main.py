N = int(input())
tree = {}
for _ in range(N):
    n, l, r = input().split()
    tree[n] = (l, r)

def traverse(n, order):
    if n == '.':
        return []
    sub = [traverse(tree[n][0], order), [n], traverse(tree[n][1], order)]
    return sub[order[0]] + sub[order[1]] + sub[order[2]]

print(''.join(traverse('A', (1, 0, 2))))
print(''.join(traverse('A', (0, 1, 2))))
print(''.join(traverse('A', (0, 2, 1))))
