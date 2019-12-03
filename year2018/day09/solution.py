import sys
import string
import collections


all=string.maketrans('','')
no_digits = all.translate(all, string.digits)

for dataset in range(int(sys.argv[2])):
    with open(sys.argv[1], "r") as fp:
        line = fp.readlines()[dataset]

        player_count, last_value = line.split(";")
        player_count = int(player_count.translate(all, no_digits))
        last_value = int(last_value.translate(all, no_digits))

    print(player_count, last_value)

    # dirty elves >.>
    players = collections.defaultdict(int)
    board = collections.deque([0])

    for m in range(1, last_value+1):
        if m % 23 == 0:
            board.rotate(7)
            players[m % player_count] += m + board.pop()
            board.rotate(-1)
            continue

        board.rotate(-1)
        board.append(m)

    print("HIGH SCORE:", max(players.values()))
