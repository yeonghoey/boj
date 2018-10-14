N = int(input())
B = [[int(s) for s in input().split()] for _ in range(N)]

def go(board, prev, depth):
    x = largest(board)
    if depth == 5:
        return x

    if prev and equal(board, prev):
        return x

    return max(
        go(L(board), board, depth + 1),
        go(R(board), board, depth + 1),
        go(U(board), board, depth + 1),
        go(D(board), board, depth + 1),
    )

def L(board):
    return move(board)

def R(board):
    return flip(move(flip(board)))

def U(board):
    return transpose(move(transpose(board)))

def D(board):
    return transpose(flip(move(flip(transpose(board)))))

def move(board):
    boardnext = []
    for row in board:
        rownext = []
        z = 0
        joinable = False
        for c in row:
            if c == 0:
                z += 1
            else:
                if joinable and rownext[-1] == c:
                    joinable = False
                    rownext[-1] += c
                    z += 1
                else:
                    rownext.append(c)
                    joinable = True

        for _ in range(z):
            rownext.append(0)
        boardnext.append(rownext)
    return boardnext

def flip(board):
    return [[c for c in reversed(row)] for row in board]

def transpose(board):
    return [[board[j][i] for j in range(N)] for i in range(N)]

def largest(board):
    return max(max(row) for row in board)

def equal(board1, board2):
    return all(
        all(board1[i][j] == board2[i][j] for j in range(N))
        for i in range(N)
    )

ans = go(B, None, 0)
print(ans)
