P = 1000000007
N, K = (int(s) for s in input().split())

F = [0]*(N+1)
F[0] = 1
for n in range(1, N+1):
    F[n] = (F[n-1] * n) % P

Ki = pow(F[K], P-2, P)
NKi =  pow(F[N-K], P-2, P)
print((((F[N]*Ki) % P) * NKi) % P)
