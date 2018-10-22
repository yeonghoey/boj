from random import randint
N, K = (int(term) for term in input().split())
A = [int(term) for term in input().split()]

begin = 0
end = N
k = K-1
answer = None

while True:
    length = end - begin
    if length < 1:
        break

    pi, i = randint(begin, end-1), begin
    A[begin], A[pi] = A[pi], A[begin]
    pivot = A[begin]
    for j in range(begin+1, end):
        if A[j] < pivot:
            i += 1
            A[i], A[j] = A[j], A[i]
    A[begin], A[i] = A[i], A[begin]

    if i == k:
        answer = A[i]
        break

    if i < k:
        begin, end, k = i+1, end, k-i
        continue

    if i > k:
        begin, end, k = begin, i, k
        continue

print(answer)
