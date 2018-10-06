from string import ascii_lowercase
word = input()
d = {}
for i, c in enumerate(word):
    d.setdefault(c, i)
print(' '.join(str(d.get(c, -1)) for c in ascii_lowercase))
