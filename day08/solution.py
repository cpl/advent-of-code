import sys


digits = []
with open(sys.argv[1], "r") as fp:
    c = fp.read()
    digits = [int(d) for d in c.split(' ')]

nodes = []
total = 0

index = 0
while index != len(digits):
    nodes.append([digits[index], digits[index+1]])
    index += 2


    done = False
    while not done:
        print(nodes)
        if nodes[-1][0] == 0:
            total += sum(digits[index:index+nodes[-1][1]])
            index += nodes[-1][1]
            nodes.pop()
            if len(nodes) == 0:
                break

            nodes[-1][0] -= 1
            done = True
            if nodes[-1][0] == 0:
                done = False
                continue
        done = True

    print(nodes)


print(total)