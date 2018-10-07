N = int(input())
A = [int(input()) for _ in range(N)]

def mergesort(a):
    if len(a) <= 1:
        return a
    mid = len(a) // 2
    l = mergesort(a[:mid])
    r = mergesort(a[mid:])
    merge(a, l, r)
    return a

def merge(a, l, r):
    ai, li, ri = 0, 0, 0

    while li < len(l) and ri < len(r):
        if l[li] < r[ri]:
            a[ai] = l[li]
            li += 1
        else:
            a[ai] = r[ri]
            ri += 1
        ai += 1

    while li < len(l):
        a[ai] = l[li]
        li += 1
        ai += 1

    while ri < len(r):
        a[ai] = r[ri]
        ri += 1
        ai += 1

mergesort(A)

for n in A:
    print(n)
