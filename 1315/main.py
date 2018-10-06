from collections import defaultdict
n = int(input())
words = (input() for _ in range(n))

def group(w):
    d = defaultdict(list)
    for i, c in enumerate(w):
        d[c].append(i)
    for k, idxs in d.items():
        if (max(idxs) - min(idxs) + 1) != len(idxs):
            return 0
    return 1

print(sum(group(w) for w in words))
