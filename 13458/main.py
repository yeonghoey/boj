N = int(input())
A = map(int, input().split())
B, C = map(int, input().split())
print(sum(((a-B-1) // C) + 1 if a >= B else 0 for a in A) + N)
