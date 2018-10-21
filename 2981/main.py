from math import sqrt

N = int(input())
nums = sorted(int(input()) for _ in range(N))
diff = nums[-1] - nums[0]

cands = set(elem
            for div in range(1, int(sqrt(diff))+1)
            if diff % div == 0
            for elem in (div, diff // div))

Ms = sorted(cand
            for cand in cands
            if cand > 1
            if all(num % cand == nums[0] % cand for num in nums))

print(' '.join(str(m) for m in Ms))
