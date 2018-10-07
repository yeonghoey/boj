N = int(input())

words = set(input().strip() for _ in range(N))
words = sorted(words, key=lambda w: (len(w), w))

for w in words:
    print(w)
