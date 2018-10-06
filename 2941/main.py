from itertools import tee, zip_longest

tgt = set('c= c- dz= d- lj nj s= z='.split())

def nx(it):
    next(it, None)
    return it

word = input()
it1, it2, it3 = tee(word, 3)
cands = ((a+b, a+b+c)
         for a, b, c in zip_longest(it1, nx(it2), nx(nx(it3)), fillvalue=' '))

ans = len(word) - sum(1 if two in tgt or three in tgt else 0
                      for two, three in cands)
print(ans)
