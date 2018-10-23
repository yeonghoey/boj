from itertools import cycle

r1, c1, r2, c2 = (int(term) for term in input().split())

n = max(abs(r1), abs(r2), abs(c1), abs(c2))
l = 2*n + 1

grid = [[''] * (c2-c1+1) for i in range(r2-r1+1)]
y, x = n, n
directions = cycle([(0, 1), (-1, 0), (0, -1), (1, 0)])
direction = next(directions)
value = 1
width, lastwidth, incwidth = 1, 1, False
while value <= l*l:
    yy, xx = y-n, x-n
    if r1 <= yy <= r2 and c1 <= xx <= c2:
        grid[yy-r1][xx-c1] = str(value)
    dy, dx = direction
    y += dy
    x += dx
    value += 1
    width -= 1
    if width == 0:
        if incwidth:
            lastwidth += 1
        incwidth = not incwidth
        width = lastwidth
        direction = next(directions)

digit = max(len(cell) for row in grid for cell in row)
for row in grid:
    print(' '.join('{:>{digit}}'.format(cell, digit=digit) for cell in row))
