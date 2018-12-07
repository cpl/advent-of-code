from collections import defaultdict


def mdst(x, xp, y, yp):
    return abs(x - xp) + abs(y - yp)


points = set()
max_x = 0
max_y = 0


with open("day06/input.txt", "r") as fp:
    for line in fp.readlines():
        x, y = map(int, line.split(", "))
        max_x = max(x, max_x)
        max_y = max(y, max_y)
        points.add((x, y))


idx2pnt = {idx: p for idx, p in enumerate(points)}
areas = defaultdict(int)
infidx = set()

total = 0
for xp in range(max_x + 1):
    for yp in range(max_y + 1):
        dists = sorted([ (mdst(x, xp, y, yp), idx) for idx, (x, y) in idx2pnt.items()])

        total += int(sum(mdst(x, xp, y, yp) for x, y in points) < 10000)

        if len(dists) == 1 or dists[0][0] != dists[1][0]:
            idx = dists[0][1]
            areas[idx] += 1

            if xp == 0 or xp == max_x or yp == 0 or yp == max_y:
                infidx.add(idx)


print("DISTANCES")
print(dists)

print("AREAS")
print(sorted([(idx, area) for idx, area in areas.items()]))

print("INFINITE POINTS")
print(["%d: {%d, %d}" % (idx, point[0], point[1]) for idx, point in idx2pnt.items() if idx in infidx])

print("FINITE POINTS")
print(["%d: {%d, %d}" % (idx, point[0], point[1]) for idx, point in idx2pnt.items() if idx not in infidx])

print("MAX AREA OF FINITE POINTS")
print(max([area for idx, area in areas.items() if idx not in infidx]))

print("TOTAL")
print(total)
