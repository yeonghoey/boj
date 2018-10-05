raw = (int(input()) for _ in range(5))
c40 = (s if s > 40 else 40 for s in raw)
print(sum(c40) // 5)
