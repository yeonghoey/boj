n, m = (int(s) for s in input().split())
fac = [1]
for a in range(1, n+1):
    fac.append(fac[-1]*a)
print(fac[n] // fac[m] // fac[n-m])
