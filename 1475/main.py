from collections import Counter
counter = Counter(input())
need = (counter['6'] + counter['9'] + 1) // 2
counter['6'] = need
counter['9'] = need
_, count = counter.most_common(1)[0]
print(count)
