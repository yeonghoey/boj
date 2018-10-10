from itertools import combinations as C
N,M=map(int, input().split())
l,s,v={},[],[]
for i in range(N):
	r=list(map(int, input().split()))
	for j in range(M):
		l[(i, j)]=r[j]
		if r[j]==0:
			s.append((i,j))
		if r[j]==2:
			v.append((i,j))
be=0
for a,b,c in C(s, 3):
	l[a],l[b],l[c]=1,1,1
	st=v[:]
	vi=set(v)
	os=[(1,0),(0,1),(-1,0),(0,-1)]
	sa,ab=len(s)-3,False
	while st:
		i,j=st.pop()
		for p in [(i+y,j+x) for y,x in os]:
			if p not in vi and l.get(p)==0:
				vi.add(p)
				st.append(p)
				sa-=1
				if sa<be:
					ab=True
		if ab:
			break
	if sa>be:
		be=sa
	l[a],l[b],l[c]=0,0,0
print(be)
