import sys
import string

all=string.maketrans('','')
no_digits = all.translate(all, string.digits)

for dataset in range(int(sys.argv[2])):
    with open(sys.argv[1], "r") as fp:
        line = fp.readlines()[dataset]

        player_count, last_value = line.split(";")
        player_count = int(player_count.translate(all, no_digits))
        last_value = int(last_value.translate(all, no_digits))

    print(player_count, last_value)


    current = 4
    players = [0 for _ in range(player_count)]
    board = [0, 4, 2, 1, 3]

    # print(players)
    # print(board)

    pturn = 4
    for m in range(5, last_value+1):

        if m % 23 == 0:
            idx = board.index(current)-7
            if idx < 0:
                idx += len(board)

            players[pturn] += m
            players[pturn] += board[idx]
            board.pop(idx)
            current = board[idx]
            pturn = (pturn+1)%player_count
            continue

        # normal insert
        idx = board.index(current)+2
        if idx > len(board):
            idx = idx % len(board)
        board.insert(idx, m)
        current = m

        # debug
        # print(m)
        # print(board)

        # update player turn
        pturn = (pturn+1)%player_count

    print(players)
    print("HIGH SCORE:", max(players))
