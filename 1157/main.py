from collections import Counter
word = input().upper()
c = Counter(word)
most = c.most_common(2)
answer = most[0][0]
if len(most) == 2 and most[0][1] == most[1][1]:
    answer = '?'
print(answer)
