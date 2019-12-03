import sys
import pprint


WORKER_COUNT = 5
workers = [[0, ''] for _ in range(WORKER_COUNT)]
total_time = 0

task_deps = {}
solution = ''


def passtime():
    global total_time
    total_time+=1
    print("WORKERS", workers)
    print(total_time, "TIME++")
    for wid in range(WORKER_COUNT):
        if workers[wid][0] > 0:
            workers[wid][0] -= 1
        if workers[wid][0] == 0 and workers[wid][1] != '':
            ft = workers[wid][1]
            workers[wid][1] = ''
            done(ft)

def done(source):
    global solution
    global task_deps

    del task_deps[source]
    solution += source

    print(source, "ENTRY")
    pprint.pprint(task_deps)

    for task in task_deps.keys():
        if task_deps.get(task) is not None and source in task_deps[task]:
            for index, val in enumerate(task_deps[task]):
                if val == source:
                    task_deps[task].pop(index)
                    print("REMOVING", val, "FROM", task, task_deps[task])


with open(sys.argv[1], "r") as fp:
    for line in fp.readlines():
        prv = line[5]
        nxt = line[36]

        if task_deps.get(prv) is None:
            task_deps[prv] = []
        if task_deps.get(nxt) is None:
            task_deps[nxt] = []

        task_deps[nxt].append(prv)


print(task_deps)

while task_deps != {}:
    for key in task_deps.keys():
        if task_deps.get(key) is not None and len(task_deps[key]) == 0:

            skip_this = False
            for wid in range(WORKER_COUNT):
                if workers[wid][1] == key:
                    skip_this = True

            if skip_this:
                continue

            # find free worker
            for wid in range(WORKER_COUNT):
                if workers[wid][0] == 0:
                    workers[wid] = [ord(key) - ord('A') + 61, key]
                    print("WORKER", wid, key)
                    found_worker = True
                    break

    passtime()



print(task_deps)
print("Solution:", solution)
print("Total time:", total_time)
