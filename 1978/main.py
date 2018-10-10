MAX_N = 1001
isprime = [True] * MAX_N
isprime[0] = False
isprime[1] = False

for i in range(2, MAX_N):
    if not isprime[i]:
        continue

    for j in range(i+i, MAX_N, i):
        isprime[j] = False

_ = int(input())
numbers = map(int, input().split())
answer = sum(1 for n in numbers if isprime[n])
print(answer)
