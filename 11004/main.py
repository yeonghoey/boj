import sys

N, K = (int(term) for term in sys.stdin.readline().split())
A = sys.stdin.readline().split()
for i in range(N):
    A[i] = int(A[i])

K -= 1

def partition(begin, end):
    lo, pi = begin, end-1
    pivot = A[pi]
    for i in range(begin, end-1):
        if A[i] <= pivot:
            A[lo], A[i] = A[i], A[lo]
            lo += 1
    A[lo], A[pi] = A[pi], A[lo]
    return lo


def f(begin, end):
    lo = partition(begin, end)

    if lo < K:
        return f(lo+1, end)

    if lo > K:
        return f(begin, lo)

    return A[K]

print(f(0, N))
