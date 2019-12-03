import re
import time


class Entry(object):

    def __init__(self, s):
        self.date_str = s[1:17]
        self.date = time.strptime(self.date_str, '%Y-%m-%d %H:%M')
        self.s = s[19:]

    def __str__(self):
        return self.date_str

    def __eq__(self, value):
        return self.date.__eq__(value.date)

    def __gt__(self, value):
        return self.date.__gt__(value.date)





if __name__ == '__main__':
    guards = {}

    # Read entries
    entries = []
    with open("input.txt", "r") as fp:
        for line in fp.readlines():
            entries.append(Entry(line.strip()))

    # Sort entries
    entries.sort()

    # Save sorted entries
    with open("input.sorted.txt", "w+") as fp:
        for entry in entries:
            fp.write("[%s] %s\n" % (entry.date_str, entry.s))

    # Find sleepy
    gid = 0
    fsm = 0
    for entry in entries:

        if entry.s.startswith("Guard"):
            gid = int(re.search(r'\d+', entry.s).group())

            if guards.get(gid) is None:
                guards[gid] = [0 for _ in range(61)]
        elif entry.s == "falls asleep":
            fsm = entry.date.tm_min
        elif entry.s == "wakes up":
            guards[gid][0] += entry.date.tm_min - fsm
            for i in range(fsm, entry.date.tm_min):
                guards[gid][i] += 1

    max_mins = 0
    max_gid = 0
    for k, v in guards.items():
        if v[0] > max_mins:
            max_mins = v[0]
            max_gid = k

    print("PART 1")
    print("GUARD ID: %d" % max_gid)

    max_min = 0
    max_mini = 0
    for index, m in enumerate(guards[max_gid][1:]):
        if m > max_min:
            max_min = m
            max_mini = index

    print("MINUTE: %d" % (max_mini+1))
    print("ANS: %d" % ((max_mini+1)*max_gid))


    max_min = 0
    max_mini = 0
    max_gid = 0
    for k, v in guards.items():
        for index, m in enumerate(v[1:]):
            if m > max_min:
                max_min = m
                max_gid = k
                max_mini = index

    print("PART 2")
    print("GUARD ID: %d" % max_gid)
    print("MINUTE: %d" % (max_mini+1))
    print("ANS: %d" % (max_gid*(max_mini+1)))
