n = int(input())
best = [0]
for _ in range(n):
    row = [int(term) for term in input().split()]
    best.append(best[-1]) # best[len(best)] == best[len(best)-1]
    best.append(best[0])  # best[0-1] == best[0]
    best = [max(best[i-1], best[i]) + n for i, n in enumerate(row)]
print(max(best))
