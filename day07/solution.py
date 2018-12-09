import sys
import pprint


task_deps = {}
solution = ''

def done(source):
    global solution
    global task_deps

    del task_deps[source]
    solution += source

    print(source, "ENTRY")
    pprint.pprint(task_deps)

    for task in sorted(task_deps.keys()):
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
    for key in sorted(task_deps.keys()):
        if task_deps.get(key) is not None and len(task_deps[key]) == 0:
            done(key)
            break

print(task_deps)
print("Solution:", solution)
