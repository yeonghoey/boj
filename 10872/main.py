N = int(input())
fac = [1]
for n in range(1, N+1):
    fac.append(fac[-1]*n)
print(fac[N])
