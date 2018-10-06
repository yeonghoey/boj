def gen():
    k = 2
    while True:
        for l in range(1, k):
            yield ((k-l, l) if k % 2 == 0 else
                   (l, k-l))
        k += 1

x = int(input())
for _, (de, nu) in zip(range(x), gen()):
    pass
print('%s/%s' % (de, nu))
