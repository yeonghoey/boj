chrs = 'ABC DEF GHI JKL MNO PQRS TUV WXYZ'.split()
secs = range(3, 11)
d = {c: s for cs, s in zip(chrs, secs) for c in cs}
print(sum(d[c] for c in input()))
