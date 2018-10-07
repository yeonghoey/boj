from collections import Counter
N = int(input())

numbers = [int(input()) for _ in range(N)]

mean = round(sum(numbers) / N)
print(mean)

median = sorted(numbers)[N//2]
print(median)

counter = Counter(numbers)
maxfreq = counter.most_common(1)[0][1]
mosts = sorted(k for k, v in counter.items() if v == maxfreq)
most_common = (mosts[0] if len(mosts) == 1 else mosts[1])
print(most_common)

range_ = max(numbers) - min(numbers)
print(range_)
