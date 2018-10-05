t = int(input())
for _ in range(t):
    s = input()
    score, acc = 0, 1
    for c in s:
        if c == 'O':
            score += acc
            acc += 1
        else:
            acc = 1
    print(score)
